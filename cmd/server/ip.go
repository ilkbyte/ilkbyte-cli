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

// iplistCmd represents the iplist command
var IpListCmd = &cobra.Command{
	Use:   "iplist",
	Short: "Get your server's ip addresses",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetServerIPList(name)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Address", "Reverse", "ACL List", "Usage"}

		v4Data := [][]string{}
		for _, v := range resp.Data.IPV4 {
			val := []string{v.Address, v.Reverse, v.ACLList, v.Usage}
			v4Data = append(v4Data, val)
		}

		v6Data := [][]string{}
		for _, v := range resp.Data.IPV6 {
			val := []string{v.Address, v.Reverse, v.ACLList, v.Usage}
			v6Data = append(v6Data, val)
		}

		fmt.Println("\nIPv4 Information")
		table.Create(header, v4Data)
		fmt.Println("\nIPv6 Information")
		table.Create(header, v6Data)
	},
}

// iplogsCmd represents the iplogs command
var IpLogsCmd = &cobra.Command{
	Use:   "iplogs",
	Short: "Get your server's ip logs",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetServerIPLogs(name)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"IP Prefix", "Is Person", "Is Log", "Log File", "Rule In", "Rule Out", "Rule Type"}
		data := [][]string{}
		for _, v := range resp.Data {
			val := []string{v.IPPrefix, fmt.Sprintf("%v", v.IsPerson), fmt.Sprintf("%v", v.IsLog), v.LogFile, fmt.Sprintf("%v", v.RuleIn), fmt.Sprintf("%v", v.RuleOut), v.RuleType}
			data = append(data, val)
		}

		table.Create(header, data)
	},
}

func init() {
	IpListCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	IpListCmd.MarkFlagRequired("name")

	IpLogsCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	IpLogsCmd.MarkFlagRequired("name")
}
