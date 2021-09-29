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
package server

import (
	"fmt"
	"log"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/ilkbyte/ilkbyte-cli/utils/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// detailCmd represents the detail command
var DetailCmd = &cobra.Command{
	Use:   "detail",
	Short: "Server Details",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetServerDetails(name)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"IPv4", "IPv6", "Bandwidth Limit", "Bandwidth Usage", "Service", "Status"}
		data := [][]string{
			[]string{resp.Data.IPV4, resp.Data.IPV6, fmt.Sprintf("%v", resp.Data.BandwidthLimit), fmt.Sprintf("%v", resp.Data.BandwidthUsage), resp.Data.Service, resp.Data.Status},
		}

		table.Create(header, data)
	},
}

func init() {
	DetailCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	DetailCmd.MarkFlagRequired("name")
}
