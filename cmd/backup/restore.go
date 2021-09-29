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
package backup

import (
	"fmt"
	"log"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/ilkbyte/ilkbyte-cli/utils/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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
	BackupCmd.AddCommand(restoreBackupCmd)
	restoreBackupCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	restoreBackupCmd.MarkFlagRequired("name")

	restoreBackupCmd.Flags().StringVarP(&backupName, "backup-name", "b", "", "restore backup name")
	restoreBackupCmd.MarkFlagRequired("backup-name")
}
