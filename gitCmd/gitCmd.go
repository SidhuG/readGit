//*****************************************************************
// Package: gitCmd
// Purpose: Performs git operations
// Author: SidhuG
//*******************************************************************

package gitCmd

import (
	"fmt"
	//homedir "github.com/mitchellh/go-homedir"
	//"gopkg.in/libgit2/git2go.v25"
	"os"
	"path/filepath"
	homedir "github.com/mitchellh/go-homedir"

	//"strconv"
	//"reflect"
)

type RepoStruct struct {
	GitUrl, GitUser, SshId, ProjectRepo, GitBranch, GitTag string
}

func emptydir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckOutRepo(rep RepoStruct) (dirPath string, err error) {

	//1. Make tmp dir
	tmppath := filepath.Join(".", "tmp")
	if _, err := os.Stat(tmppath); os.IsNotExist(err) {
		os.Mkdir(tmppath, 0755)
	}

	//2. Clean tmp directory if needed
	if err := emptydir(tmppath); err != nil {
		fmt.Println("FATAL: can not empty existing directory")
	}

	//3. git clone
	//git.CloneOption{}
	//repo, err := git.Clone("git://github.com/gopheracademy/gopheracademy-web.git", "web", &git.CloneOptions{})
	//if err != nil {
	//	panic(err)
	//}
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sshKeyFile := filepath.Join(home, rep.SshId)
	setSSHCredentials(sshKeyFile)
	checkoutBranch("git@"+rep.GitUrl+":"+rep.GitUser+"/"+rep.ProjectRepo, rep.GitBranch)

	//4. return the path to where repo has been cloned

	return tmppath, err

}
