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
package domain

import (
	"fmt"
	"log"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/ilkbyte/ilkbyte-cli/utils/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the domain command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get your all domains",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetDomains(page)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Name", "NS-Master", "NS-Slave"}
		data := [][]string{}
		for _, v := range resp.Data.DomainList {
			val := []string{v.Domain, v.Nameserver.Master, v.Nameserver.Slave}
			data = append(data, val)
		}

		table.Create(header, data)

		fmt.Println("Showing", resp.Data.Pagination.CurrentPage, "of", resp.Data.Pagination.TotalPage, "pages")
	},
}

func init() {
	DomainCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&page, "page", "p", 1, "paginate your domains")
}
