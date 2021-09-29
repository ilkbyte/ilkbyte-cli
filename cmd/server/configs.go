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

// configsCmd represents the configs command
var ConfigsCmd = &cobra.Command{
	Use:   "configs",
	Short: "Get configs for create server",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetServerConfig()
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		appHeader := []string{"Name", "Code", "System"}
		appData := [][]string{}
		for _, v := range resp.Data.Application {
			val := []string{v.Name, fmt.Sprintf("%v", v.Code), v.System}
			appData = append(appData, val)
		}

		osHeader := []string{"Name", "Code", "Version"}
		osData := [][]string{}
		for _, v := range resp.Data.OperatingSystem {
			val := []string{v.Name, fmt.Sprintf("%v", v.Code), v.Version}
			osData = append(osData, val)
		}

		packHeader := []string{"Name", "Code", "Features", "Price"}
		packData := [][]string{}
		for _, v := range resp.Data.Package {
			val := []string{v.Name, fmt.Sprintf("%v", v.Code), v.Features, v.Price}
			packData = append(packData, val)
		}

		fmt.Println("\nApplication Parameters")
		table.Create(appHeader, appData)
		fmt.Println("\nOS Parameters")
		table.Create(osHeader, osData)
		fmt.Println("\nPackage Parameters")
		table.Create(packHeader, packData)
	},
}
