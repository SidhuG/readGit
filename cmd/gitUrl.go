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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	//"github.com/SidhuG/readGit/gitCmd"
)

var git_url string
var git_tag string
var FQDN_list string

// gitUrlCmd represents the gitUrl command
var gitUrlCmd = &cobra.Command{
	Use:   "gitUrl",
	Short: "remote repo url",
	Long:  `Checks out remote git repo`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitUrl command used, values in config file will be overridden with following")
		fmt.Println("--- Git URL: ", git_url)
		fmt.Println("--- Git Tag: ", git_tag)
		fmt.Println("--- FQDNs: ", FQDN_list)
	},
}

func init() {
	RootCmd.AddCommand(gitUrlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	//git url where configuration files are located
	gitUrlCmd.PersistentFlags().StringVarP(&git_url, "url", "u", "", "git remote url")

	//git tag to release
	gitUrlCmd.PersistentFlags().StringVarP(&git_tag, "tag", "t", "", "tag to use")

	//list of FQDN to apply this configuration change
	gitUrlCmd.PersistentFlags().StringVarP(&FQDN_list, "fqdns", "f", "", "FQDNs to apply this configuration")

	viper.BindPFlag("url", gitUrlCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("tag", gitUrlCmd.PersistentFlags().Lookup("tag"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gitUrlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
