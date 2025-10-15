/*
Copyright © 2021 Umut Aktepe <umtaktpe@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ilkbyte/ilkbyte-cli/cmd/account"
	"github.com/ilkbyte/ilkbyte-cli/cmd/backup"
	"github.com/ilkbyte/ilkbyte-cli/cmd/domain"
	"github.com/ilkbyte/ilkbyte-cli/cmd/server"
	"github.com/ilkbyte/ilkbyte-cli/cmd/snapshot"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ilkbyte",
	Short: "Manage your servers via terminal",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Account commands
	rootCmd.AddCommand(account.AccountCmd)

	// Snapshot commands
	rootCmd.AddCommand(snapshot.SnapshotCmd)

	// Backup commands
	rootCmd.AddCommand(backup.BackupCmd)

	// Server commands
	rootCmd.AddCommand(server.ListCmd)
	rootCmd.AddCommand(server.ConfigsCmd)
	rootCmd.AddCommand(server.CreateCmd)
	rootCmd.AddCommand(server.DetailCmd)
	rootCmd.AddCommand(server.PowerCmd)
	rootCmd.AddCommand(server.IpListCmd)
	rootCmd.AddCommand(server.IpLogsCmd)
	rootCmd.AddCommand(server.RdnsCmd)
	rootCmd.AddCommand(server.DeleteCmd)

	// Domain commands
	rootCmd.AddCommand(domain.DomainCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if _, err := os.Stat(home + "/.ilkbyte-cli.yaml"); os.IsNotExist(err) {
		_, err := os.Create(home + "/.ilkbyte-cli.yaml")
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		type conf struct {
			Ilkbyte struct {
				AccessKey string `yaml:"access_key"`
				SecretKey string `yaml:"secret_key"`
			} `yaml:"ilkbyte"`
		}

		config := &conf{}

		data, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if err := ioutil.WriteFile(home+"/.ilkbyte-cli.yaml", data, 0); err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.

		// Search config in home directory with name ".ilkbyte-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ilkbyte-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if len(viper.GetString("ilkbyte.access_key")) == 0 || len(viper.GetString("ilkbyte.secret_key")) == 0 {
			fmt.Println("Access key or secret key cannot be left blank.")
			fmt.Println("Please insert keys yaml file this path: $HOME/.ilkbyte-cli.yaml")
			os.Exit(1)
		}
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
