// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	git "github.com/SidhuG/readGit/gitCmd"
	//hiera "github.com/SidhuG/readGit/parseHiera"
	fqdn "github.com/SidhuG/readGit/verifyFQDNs"

	//kv "github.com/SidhuG/readGit/kvEndpoint"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var prjctName string
var mapGitConfig = make([]interface{}, 2)
var verifiedListCh = make([]<-chan fqdn.VerifyStatus, 0)
var status fqdn.VerifyStatus

//var gitConfigURL map[interface{}]interface{} = make(map[interface{}]interface{})

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "readGit",
	Short: "Reads git repo for config in YAML files and creates k/v endpoints",
	Long: `Checkouts Git repo at specific tag,parses yaml configuration files in the git repo,
	and creates k/v endpoints in consul/etcd
	command options can be either specifid in command line or in config file in yaml format.
	If command 'gitUrl' is specified along with all required flags then config file is ignored`,

	// This is where all the action is:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root cmd called, using config file: ", viper.ConfigFileUsed())
		viper.Debug()
		fmt.Print("Found keys in viper: ", viper.AllKeys())
		fmt.Println()

		//Extract git url, git tag and git user ssh id file path
		mapGitConfig = viper.Get("git_config").([]interface{})
		gitConfigURL := mapGitConfig[0].(map[interface{}]interface{})
		gitConfigProject := mapGitConfig[1].(map[interface{}]interface{})
		gitConfigTag := mapGitConfig[2].(map[interface{}]interface{})
		gitConfigUser := mapGitConfig[3].(map[interface{}]interface{})
		gitUserID := mapGitConfig[4].(map[interface{}]interface{})
		gitConfigBranch := mapGitConfig[5].(map[interface{}]interface{})

		fmt.Println("Using git url: ", gitConfigURL["url"])
		fmt.Println("Using git tag: ", gitConfigTag["tag"])
		fmt.Println("Using git user id: ", gitUserID["ssh_id"])
		fmt.Println("Using branch: ", gitConfigBranch["branch"])

		//Extract List of FQDNs to apply changes to
		mapFQDNList := viper.Get("fqdn_list").([]interface{})
		fmt.Println("To update hosts: ", mapFQDNList)
		//fmt.Println("Found first FQDN : " , mapFQDNList[0])

		//Checkout git repo
		rep := git.RepoStruct{
			GitURL:      gitConfigURL["url"].(string),
			GitBranch:   gitConfigBranch["branch"].(string),
			GitTag:      gitConfigTag["tag"].(string),
			SSHId:       gitUserID["ssh_id"].(string),
			GitUser:     gitConfigUser["user"].(string),
			ProjectRepo: gitConfigProject["projectrepo"].(string)}

		dirPath, err := git.CheckOutRepo(rep)
		if err != nil {
			fmt.Println("Error: Could not checkout from git!")
		} else {
			fmt.Println("Git repo cloned at: ", dirPath)
		}

		var verifyReturn <-chan fqdn.VerifyStatus
		if fqdn.LoadYaml(dirPath) == 0 {
			// traverse FQDN List
			for _, hostname := range mapFQDNList {
				//fmt.Println("Constructing endpoints for FQDN: ", hostname)
				verifyReturn = fqdn.Verify(hostname.(string))
				verifiedListCh = append(verifiedListCh, verifyReturn)
			}
		}

		for _, verified := range verifiedListCh {
			status = <-verified
			fmt.Println("Verified: ", status.Hostname, status.Status)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default search path is $HOME/.readGit.yaml, current working )")

	//RootCmd.PersistentFlags().StringVar(&prjctName,"projectname", "", "to refer a single project name if multiple git repos are specifed in the config but only 1 repo needs to be parsed")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//viper.BindPFlag("projectname", gitUrlCmd.PersistentFlags().Lookup("projectname"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".readGit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.SetConfigName("readGit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		//viper.Get("ProjectName")
	} else {
		fmt.Println("Config file Read error ")
	}
}
