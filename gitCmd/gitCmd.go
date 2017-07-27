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
	homedir "github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
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
	tmppath := filepath.Join("/tmp/", "readGit")
	if _, err := os.Stat(tmppath); os.IsNotExist(err) {
		os.Mkdir(tmppath, 0755)
	}

	//2. Clean tmp directory if needed
	if err := emptydir(tmppath); err != nil {
		fmt.Println("FATAL: can not empty existing directory")
	}

	//3. git clone
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sshKeyFile := filepath.Join(home, rep.SshId)
	setSSHCredentials(sshKeyFile)
	retpath, _ := checkoutBranch("git@"+rep.GitUrl+":"+rep.GitUser+"/"+rep.ProjectRepo, rep.GitBranch, rep.GitTag)
	if retpath != tmppath {
		log.Println("ERROR: could not checkout git repo at specific branch/tag")
	}
	//4. return the path to where repo has been cloned

	return retpath, err

}
