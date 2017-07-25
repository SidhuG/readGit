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
	hostname string,
	verifyStatus bool
}

//
func verify() <- chan verifyStatus{

}