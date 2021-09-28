/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"ilkbyte-cli/utils/client"
	"ilkbyte-cli/utils/table"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Manage backup operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

// backupListCmd represents the backup list command
var backupListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get backup list",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetAllBackup(name)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Name", "Amount", "File Size", "File Hash", "Is Locked", "Backup Time"}
		data := [][]string{
			[]string{resp.Data.Backup.Name, resp.Data.Amount, resp.Data.Backup.FileSize, resp.Data.Backup.FileHash, fmt.Sprintf("%v", resp.Data.Backup.IsLocked), resp.Data.Backup.BackupTime.String()},
		}

		table.Create(header, data)
	},
}

// backupListCmd represents the backup restore command
var restoreBackupCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore your backup",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.BackupRestore(name, backupName)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Name", "File Size", "File Hash", "Is Locked", "Backup Time"}
		data := [][]string{
			[]string{resp.Data.Name, resp.Data.FileSize, resp.Data.FileHash, fmt.Sprintf("%v", resp.Data.IsLocked), resp.Data.BackupTime.String()},
		}

		table.Create(header, data)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	backupCmd.AddCommand(backupListCmd)
	backupListCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	backupListCmd.MarkFlagRequired("name")

	backupCmd.AddCommand(restoreBackupCmd)
	restoreBackupCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	restoreBackupCmd.MarkFlagRequired("name")

	restoreBackupCmd.Flags().StringVarP(&backupName, "backup-name", "b", "", "restore backup name")
	restoreBackupCmd.MarkFlagRequired("backup-name")
}
