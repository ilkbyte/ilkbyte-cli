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
	"github.com/ilkbyte/ilkbyte-cli/utils/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// snapshotListCmd represents the snapshot list command
var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get snapshot list",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetAllSnapshots(name)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Name", "Current", "State", "Location", "Parent", "Children", "Descendants", "Metadata", "Date"}
		data := [][]string{}
		for _, v := range resp.Data.Snapshots {
			val := []string{v.Name, fmt.Sprintf("%v", v.Current), v.State, v.Location, v.Parent, fmt.Sprintf("%v", v.Children), fmt.Sprintf("%v", v.Descendants), fmt.Sprintf("%v", v.Metadata), v.Date.String()}
			data = append(data, val)
		}

		table.Create(header, data)
	},
}

func init() {
	SnapshotCmd.AddCommand(snapshotListCmd)
	snapshotListCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	snapshotListCmd.MarkFlagRequired("name")
}
