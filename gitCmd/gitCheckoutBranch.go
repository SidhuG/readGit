// Package gitCmd Checkout git branch at specific tags
package gitCmd

import (
	//"errors"
	//"github.com/libgit2/git2go"
	//"errors"
	"fmt"
	"gopkg.in/libgit2/git2go.v25"
	"log"
	"strings"
)

var gitSshid string

// MyError is an error implementation that includes a time and message.
type gitCmdError struct {
	What string
}

func (e gitCmdError) Error() string {
	return fmt.Sprintf("FATAL: %v", e.What)
}

func credentialsCallback(url string, username string, allowedTypes git.CredType) (git.ErrorCode, *git.Cred) {
	//ret, cred := git.NewCredSshKey("git", gitSshid+".pub", gitSshid, "")
	ret, cred := git.NewCredSshKeyFromAgent(username)
	return git.ErrorCode(ret), &cred
}

// Made this one just return 0 during troubleshooting...
func certificateCheckCallback(cert *git.Certificate, valid bool, hostname string) git.ErrorCode {
	return 0
}

func setSSHCredentials(sshid string) int {
	gitSshid = sshid
	log.Println("Setting key file to : ", gitSshid)
	return 0
}

func checkoutBranch(gitURL string, branchName string, tagToUse string) (string, error) {

	var tmpDirPath = "/tmp/readGit"

	cbs := git.RemoteCallbacks{
		CredentialsCallback:      credentialsCallback,
		CertificateCheckCallback: certificateCheckCallback,
	}

	cloneOptions := &git.CloneOptions{}
	cloneOptions.FetchOptions = &git.FetchOptions{DownloadTags: git.DownloadTagsAll}
	cloneOptions.CheckoutOpts = &git.CheckoutOpts{Strategy: git.CheckoutSafe | git.CheckoutRecreateMissing | git.CheckoutAllowConflicts | git.CheckoutUseTheirs}
	cloneOptions.CheckoutOpts.Strategy = 1 //Otherwise it is dry run. Nothing really clones
	cloneOptions.FetchOptions.RemoteCallbacks = cbs

	fmt.Println("About to clone: ", gitURL)
	repo, err := git.Clone(gitURL, tmpDirPath, cloneOptions)
	if err != nil {
		log.Panic(err)
		//log.Println("FATAL: could not clone")
		//return err
	}

	checkoutOpts := &git.CheckoutOpts{
		Strategy: git.CheckoutSafe | git.CheckoutRecreateMissing | git.CheckoutAllowConflicts | git.CheckoutUseTheirs,
	}

	//Parse tags
	iter, err := repo.NewReferenceIterator()
	var tagid *git.Oid
	var tagObjID *git.Oid
	var tagName string
	ref, err := iter.Next()
	for err == nil {
		if ref.IsTag() {
			fmt.Println(ref.Name())
			tagName = strings.TrimPrefix(ref.Name(), "refs/tags/")
			if tagName == tagToUse {
				tagid = ref.Target()
				tag, err := repo.LookupTag(tagid)
				if err != nil {
					fmt.Println("Could not look up tag Id")
				}
				tagObjID = tag.TargetId()
				break
			}
		}
		ref, err = iter.Next()
	}
	if tagName != tagToUse {
		log.Println("FATAL: Could not find requested tag: ", tagToUse)
		return "", gitCmdError{"Could not find requested tag"}
	}

	//Getting the reference for the remote branch
	// remoteBranch, err := repo.References.Lookup("refs/remotes/origin/" + branchName)
	remoteBranch, err := repo.LookupBranch("origin/"+branchName, git.BranchRemote)
	if err != nil {
		log.Panic(err)
		fmt.Println("Failed to find remote branch: " + branchName)
		return "", gitCmdError{"Failed to find remote branch:"}
	}
	defer remoteBranch.Free()

	//Find commit for the tag
	headCommit, err := repo.LookupCommit(tagObjID)
	if err != nil {
		panic(err)
	}
	//return nil
	localBranchName := "v" + tagToUse
	//Create a local branch at specified tag
	localBranch, err := repo.CreateBranch(localBranchName, headCommit, false)
	if err != nil {
		fmt.Println("Failed to create local branch: " + localBranchName)
		return "", gitCmdError{"Failed to create local branch."}
	}
	defer localBranch.Free()

	// Getting the tree for the branch
	localCommit, err := repo.LookupCommit(localBranch.Target())
	if err != nil {
		log.Print("Failed to lookup for commit in local branch " + localBranchName)
		return "", gitCmdError{"Failed to lookup for commit in local branch"}
	}
	defer localCommit.Free()

	tree, err := repo.LookupTree(localCommit.TreeId())
	if err != nil {
		log.Print("Failed to lookup for tree " + localBranchName)
		return "", gitCmdError{"Failed to lookup for tree"}
	}
	defer tree.Free()

	// Checkout the tree
	err = repo.CheckoutTree(tree, checkoutOpts)
	if err != nil {
		log.Print("Failed to checkout tree " + localBranchName)
		return "", gitCmdError{"Failed to checkout tree"}
	}

	// Setting the Head to point to our branch
	repo.SetHead("refs/heads/" + localBranchName)

	return tmpDirPath, gitCmdError{""}
}
