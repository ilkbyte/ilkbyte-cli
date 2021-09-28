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
	"os"

	"github.com/ilkbyte/ilkbyte-cli/utils/table"

	"github.com/ilkbyte/ilkbyte-cli/utils/client"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Get your account information",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get details your account information",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetAccountInfo()
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		var s string
		if resp.Data.Status {
			s = "Active"
		} else {
			s = "Passive"
		}

		header := []string{"Name", "Group", "Balance", "Status"}
		data := [][]string{
			[]string{resp.Data.Name, resp.Data.Group, fmt.Sprintf("%v", resp.Data.Balance), s},
		}

		table.Create(header, data)
	},
}

// userCmd represents the users command
var userCmd = &cobra.Command{
	Use:   "users",
	Short: "Get users belonging to account",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.GetAccountUsers(page)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		header := []string{"Email", "Name", "Level", "Status", "Last Login", "Verify"}
		data := [][]string{}
		for _, v := range resp.Data.UserList {
			val := []string{v.Email, v.Firstname + " " + v.Lastname, v.Level, v.Status, v.LastLogin, fmt.Sprintf("%v", v.Verify)}
			data = append(data, val)
		}

		table.Create(header, data)

		fmt.Println("Showing", resp.Data.Pagination.CurrentPage, "of", resp.Data.Pagination.TotalPage, "pages")
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)

	accountCmd.AddCommand(infoCmd)

	accountCmd.AddCommand(userCmd)
	userCmd.Flags().IntVarP(&page, "page", "p", 1, "paginate your users")
}
