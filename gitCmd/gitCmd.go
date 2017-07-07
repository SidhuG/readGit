//*****************************************************************
// Package: gitCmd
// Purpose: Performs git operations
// Author: SidhuG
//*******************************************************************

package gitCmd

import (
	"fmt"
	"os"
	"gopkg.in/libgit2/git2go.v25"
	homedir "github.com/mitchellh/go-homedir"
	"path/filepath"
	//"strconv"
	//"reflect"
)

func CheckOutRepo(giturl string, gitbranch string, gittag string, sshid string, ) (dirPath string, err error){
	
	//1. Make tmp dir
	//2. Clean tmp directory if needed
	
	//3. git clone

	//4. return the path to where repo has been cloned


}