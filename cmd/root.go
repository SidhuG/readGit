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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var prjctName string
var map_git_config []interface{} = make([]interface{},2)
//var git_config_url map[interface{}]interface{} = make(map[interface{}]interface{})

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
		fmt.Print("Found keys in viper: ",viper.AllKeys())
		fmt.Println()

		//Extract git url and git tag
		map_git_config = viper.Get("git_config").([]interface{})
		git_config_url := map_git_config[0].(map[interface{}]interface{})
		git_config_tag := map_git_config[1].(map[interface{}]interface{})
		fmt.Println("Found git url: ",git_config_url["url"])
		fmt.Println("Found git tag: ", git_config_tag["tag"])

		//Extract List of FQDNs to apply changes to
		map_FQDN_list := viper.Get("fqdn_list").([]interface{})
		fmt.Println("Found FQDN List: ", map_FQDN_list)
		fmt.Println("Found first FQDN : " , map_FQDN_list[0])

		// traverse FQDN List
		for _, fqdn := range map_FQDN_list {
			fmt.Println("Constructing endpoints for FQDN: ", fqdn)

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
