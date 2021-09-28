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
	"log"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/ilkbyte/ilkbyte-cli/utils/table"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all your servers",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		header := []string{"Name", "OS", "IPv4", "IPv6", "Bandwidth Limit", "Bandwidth Usage", "Service", "Deleted Time"}
		data := [][]string{}
		if !all {
			resp, err := c.GetActiveServers(page)
			if err != nil {
				log.Fatalf("Active server not found!")
			}

			if !resp.Status {
				log.Fatalf("An error occured: %v\n", resp.Error)
			}

			for _, v := range resp.Data.ServerList {
				val := []string{v.Name, v.Osapp, v.Ipv4, v.Ipv6, fmt.Sprintf("%v", v.BandwidthLimit), fmt.Sprintf("%v", v.BandwidthUsage), v.Service, v.DeletedTime}
				data = append(data, val)
			}

			table.Create(header, data)
			fmt.Println("Showing", resp.Data.Pagination.CurrentPage, "of", resp.Data.Pagination.TotalPage, "pages")
		} else {
			resp, err := c.GetAllServers(page)
			if err != nil {
				log.Fatalf("An error occured: %v\n", err)
			}

			if !resp.Status {
				log.Fatalf("An error occured: %v\n", resp.Error)
			}

			for _, v := range resp.Data.ServerList {
				val := []string{v.Name, v.Osapp, v.Ipv4, v.Ipv6, fmt.Sprintf("%v", v.BandwidthLimit), fmt.Sprintf("%v", v.BandwidthUsage), v.Service, v.DeletedTime}
				data = append(data, val)
			}

			table.Create(header, data)
			fmt.Println("Showing", resp.Data.Pagination.CurrentPage, "of", resp.Data.Pagination.TotalPage, "pages")
		}
	},
}

// configsCmd represents the configs command
var configsCmd = &cobra.Command{
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

// createCmd represents the create command
var createCmd = &cobra.Command{
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

// detailCmd represents the detail command
var detailCmd = &cobra.Command{
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

// powerCmd represents the power command
var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Control your server's power status",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.PowerServer(name, status)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		fmt.Println("Server power status:", status)
	},
}

// iplistCmd represents the iplist command
var ipListCmd = &cobra.Command{
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
var ipLogsCmd = &cobra.Command{
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

// rdnsCmd represents the rdns command
var rdnsCmd = &cobra.Command{
	Use:   "rdns",
	Short: "Add new rdns record",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.ServerIPRdns(name, ip, rdns)
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
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "get your all servers")
	listCmd.Flags().IntVarP(&page, "page", "p", 1, "paginate your servers")

	rootCmd.AddCommand(configsCmd)

	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	createCmd.MarkFlagRequired("name")

	createCmd.Flags().StringVarP(&username, "username", "u", "", "username for server")
	createCmd.MarkFlagRequired("username")

	createCmd.Flags().StringVarP(&password, "password", "p", "", "password for server(must be base64)")
	createCmd.MarkFlagRequired("password")

	createCmd.Flags().StringVarP(&osid, "os-id", "", "", "operation system id for server")

	createCmd.Flags().StringVarP(&appid, "app-id", "", "", "application id for server")

	createCmd.Flags().StringVarP(&packageid, "package-id", "", "", "package id for server")
	createCmd.MarkFlagRequired("package-id")

	createCmd.Flags().StringVarP(&sshkey, "sshkey", "", "", "ssh key for server")
	createCmd.MarkFlagRequired("sshkey")

	rootCmd.AddCommand(detailCmd)
	detailCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	detailCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(powerCmd)
	powerCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	powerCmd.MarkFlagRequired("name")

	powerCmd.Flags().StringVarP(&status, "status", "s", "", "send server status. (start, shutdown, reboot, destroy)")
	powerCmd.MarkFlagRequired("status")

	rootCmd.AddCommand(ipListCmd)
	ipListCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	ipListCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(ipLogsCmd)
	ipLogsCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	ipLogsCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(rdnsCmd)
	rdnsCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	rdnsCmd.MarkFlagRequired("name")

	rdnsCmd.Flags().StringVarP(&ip, "ip", "i", "", "ip address")
	rdnsCmd.MarkFlagRequired("ip")

	rdnsCmd.Flags().StringVarP(&opt, "opt", "o", "", "operation type (create, update, delete)")
	rdnsCmd.MarkFlagRequired("opt")

	rdnsCmd.Flags().StringVarP(&rdns, "rdns", "r", "", "rdns record")
	rdnsCmd.MarkFlagRequired("rdns")
}
