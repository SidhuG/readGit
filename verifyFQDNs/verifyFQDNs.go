//*****************************************************************
// Package: gitCmd
// File: gitCheckouBranch
// Purpose: Checkout git branch at specific tags
// Author: SidhuG
//*******************************************************************

package verifyFQDNs

import (
	"fmt"
	"github.com/spf13/viper"
	"text/template"
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
	Hostname string
	Status   bool
}

var fqdnConstr string

//LoadYaml LoadYaml specified by the dir path
func LoadYaml(dirPath string) int {
	fmt.Println("LoadYaml conf file at location : ", dirPath)
	//viper.Debug()
	viper.AddConfigPath(dirPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Sprintf("Fatal error config file: %s \n", err))
	}
	fqdnConstr = viper.GetString("FQDN_CONSTRUCT")
	fmt.Println("FQDN_CONSTRUCT: ", fqdnConstr)

	pools := viper.Get("pools")
	fmt.Println("Pools: ", pools)
	return 0
}

//Verify if FWDN list contains valid FQDNs
func Verify(hostname string) <-chan VerifyStatus {
	ch := make(chan VerifyStatus)
	go func() {
		// Verify the hostname here
		//fmt.Println("Constructing endpoints for FQDN: ", hostname)
		//TODO
		//
		//Create a new template based on FQDN construct
		t := template.Must(template.New("FQDN_CONSTRUCT").Parse(fqdnConstr))
		fmt.Println("DefinedTemplates: ", t.DefinedTemplates())
		//if hostname is valid, then set verifyStatus to true, otherwise false
		vs := VerifyStatus{
			Hostname: hostname,
			Status:   true,
		}
		ch <- vs
	}()
	return ch
}
