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
package snapshot

import (
	"fmt"
	"log"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// revertSnapshotCmd represents the snapshot revert command
var revertSnapshotCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert a snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.RevertSnapshot(name, snapshotName)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		fmt.Println(resp.Message)
	},
}

func init() {
	SnapshotCmd.AddCommand(revertSnapshotCmd)
	revertSnapshotCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	revertSnapshotCmd.MarkFlagRequired("name")

	revertSnapshotCmd.Flags().StringVarP(&snapshotName, "snapshot-name", "s", "", "snapshot name")
	revertSnapshotCmd.MarkFlagRequired("snapshot-name")
}
