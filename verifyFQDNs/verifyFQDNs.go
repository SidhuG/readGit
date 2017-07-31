//*****************************************************************
// Package: gitCmd
// File: gitCheckouBranch
// Purpose: Checkout git branch at specific tags
// Author: SidhuG
//*******************************************************************

package verifyFQDNs

import (
	"fmt"
	//	"log"
)

// MyError is an error implementation that includes a time and message.
type verifyFQDNError struct {
	What string
}

func (e verifyFQDNError) Error() string {
	return fmt.Sprintf("FATAL: %v", e.What)
}

//VerifyStatus status that is returned in the channel
type VerifyStatus struct {
	hostname string
	status   bool
}

//LoadYaml LoadYaml specified by the dir path
func LoadYaml(dirPath string) int {

	return 0
}

//Verify if FWDN list contains valid FQDNs
func Verify(hostname string) <-chan VerifyStatus {
	ch := make(chan VerifyStatus)
	go func() {
		// Verify the hostname here
		fmt.Println("Constructing endpoints for FQDN: ", hostname)
		//TODO
		//
		//if hostname is valid, then set verifyStatus to true, otherwise false
		vs := VerifyStatus{
			hostname: "",
			status:   true,
		}
		ch <- vs
	}()
	return ch
}
