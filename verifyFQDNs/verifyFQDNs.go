//*****************************************************************
// Package: gitCmd
// File: gitCheckouBranch
// Purpose: Checkout git branch at specific tags
// Author: SidhuG
//*******************************************************************
package verifyFQDNs
import (
	"fmt"
	"log"
)

// MyError is an error implementation that includes a time and message.
type verifyFQDNError struct {
	What string
}

func (e verifyFQDNError) Error() string {
	return fmt.Sprintf("FATAL: %v", e.What)
}

//
type verifyStatus struct {
	hostname string
	status bool
}

func LoadYaml(dirPath string) int {

} 

//Verify if FWDN list contains valid FQDNs
func Verify(fqdns []string) <-chan verifyStatus {
	ch := make(chan verifyStatus)
	go funcVerfiy() {
		// traverse FQDN List
		for _, hostname := range fqdns {
			fmt.Println("Constructing endpoints for FQDN: ", hostname)

			//if hostname is valid, then set verifyStatus to true, otherwise false
			vs := verifyStatus {
				hostname: "",
				status: true
			}
            ch <- vs

		}
	}()
	return ch
}