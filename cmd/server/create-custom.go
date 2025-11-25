/*
Copyright © 2025 Enes Kaya <eneskaya5261@gmail.com>

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

var (
	imgURL    string
	imgSHA256 string
)

// createCustomCmd represents the create-custom command
var CreateCustomCmd = &cobra.Command{
	Use:   "create-custom",
	Short: "Create custom server with custom image",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.CreateServerWithCustomImage(name, packageid, imgURL, imgSHA256)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Name", "IPv4", "IPv6", "Service"}
		data := [][]string{
			{resp.Data.ServerInfo.Name, resp.Data.ServerInfo.IPV4, resp.Data.ServerInfo.IPV6, resp.Data.ServerInfo.Service},
		}

		table.Create(header, data)
	},
}

func init() {
	CreateCustomCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	CreateCustomCmd.MarkFlagRequired("name")

	CreateCustomCmd.Flags().StringVarP(&packageid, "package-id", "", "", "package id for server")
	CreateCustomCmd.MarkFlagRequired("package-id")

	CreateCustomCmd.Flags().StringVarP(&imgURL, "image-url", "", "", "custom image url")
	CreateCustomCmd.MarkFlagRequired("image-url")

	CreateCustomCmd.Flags().StringVarP(&imgSHA256, "image-sha256", "", "", "custom image sha256")
	CreateCustomCmd.MarkFlagRequired("image-sha256")
}
