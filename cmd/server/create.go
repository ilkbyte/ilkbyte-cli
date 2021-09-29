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
	"log"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/ilkbyte/ilkbyte-cli/utils/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create server",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.CreateServer(name, username, password, osid, appid, packageid, sshkey)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Name", "Username", "Password", "OS", "IPv4", "IPv6", "Service"}
		data := [][]string{
			[]string{resp.Data.ServerInfo.Name, resp.Data.ServerInfo.Username, resp.Data.ServerInfo.Password, resp.Data.ServerInfo.Osapp, resp.Data.ServerInfo.IPV4, resp.Data.ServerInfo.IPV6, resp.Data.ServerInfo.Service},
		}

		table.Create(header, data)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	CreateCmd.MarkFlagRequired("name")

	CreateCmd.Flags().StringVarP(&username, "username", "u", "", "username for server")
	CreateCmd.MarkFlagRequired("username")

	CreateCmd.Flags().StringVarP(&password, "password", "p", "", "password for server(must be base64)")
	CreateCmd.MarkFlagRequired("password")

	CreateCmd.Flags().StringVarP(&osid, "os-id", "", "", "operation system id for server")

	CreateCmd.Flags().StringVarP(&appid, "app-id", "", "", "application id for server")

	CreateCmd.Flags().StringVarP(&packageid, "package-id", "", "", "package id for server")
	CreateCmd.MarkFlagRequired("package-id")

	CreateCmd.Flags().StringVarP(&sshkey, "sshkey", "", "", "ssh key for server")
	CreateCmd.MarkFlagRequired("sshkey")
}
