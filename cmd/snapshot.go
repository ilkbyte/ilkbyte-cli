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

	"github.com/ilkbyte/ilkbyte-cli/utils/client"
	"github.com/ilkbyte/ilkbyte-cli/utils/table"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// snaphotCmd represents the snapshot command
var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Manage snapshot operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

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

// createSnapshotCmd represents the snapshot create command
var createSnapshotCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.CreateSnapshot(name, snapshotName)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		fmt.Println(resp.Message)
	},
}

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

// updateSnapshotCmd represents the snapshot update command
var updateSnapshotCmd = &cobra.Command{
	Use:   "update",
	Short: "Recreate a snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.UpdateSnapshot(name, snapshotName)
		if err != nil {
			log.Fatalf("An error occured: %v\n", err)
		}

		if !resp.Status {
			log.Fatalf("An error occured: %v\n", resp.Error)
		}

		fmt.Println(resp.Message)
	},
}

// deleteSnapshotCmd represents the snapshot delete command
var deleteSnapshotCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(&client.Option{
			AccessKey: viper.GetString("ilkbyte.access_key"),
			SecretKey: viper.GetString("ilkbyte.secret_key"),
		})

		resp, err := c.DeleteSnapshot(name, snapshotName)
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
	rootCmd.AddCommand(snapshotCmd)

	snapshotCmd.AddCommand(snapshotListCmd)
	snapshotListCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	snapshotListCmd.MarkFlagRequired("name")

	snapshotCmd.AddCommand(createSnapshotCmd)
	createSnapshotCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	createSnapshotCmd.MarkFlagRequired("name")

	createSnapshotCmd.Flags().StringVarP(&snapshotName, "snapshot-name", "s", "", "snapshot name")
	createSnapshotCmd.MarkFlagRequired("snapshot-name")

	snapshotCmd.AddCommand(revertSnapshotCmd)
	revertSnapshotCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	revertSnapshotCmd.MarkFlagRequired("name")

	revertSnapshotCmd.Flags().StringVarP(&snapshotName, "snapshot-name", "s", "", "snapshot name")
	revertSnapshotCmd.MarkFlagRequired("snapshot-name")

	snapshotCmd.AddCommand(deleteSnapshotCmd)
	deleteSnapshotCmd.Flags().StringVarP(&name, "name", "n", "", "server name")
	deleteSnapshotCmd.MarkFlagRequired("name")

	deleteSnapshotCmd.Flags().StringVarP(&snapshotName, "snapshot-name", "s", "", "snapshot name")
	deleteSnapshotCmd.MarkFlagRequired("snapshot-name")
}
