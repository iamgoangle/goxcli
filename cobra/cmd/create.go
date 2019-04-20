// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user in a shop",
	Long:  ``,
	Run:   create,
}

var name string
var id string

func init() {
	rootCmd.AddCommand(createCmd)

	// name
	createCmd.Flags().StringVarP(&name, "name", "n", "John Doe", "Name of user (required)")

	// id
	createCmd.Flags().StringVarP(&id, "id", "i", "", "Specify userId (required)")

	// add flag required fields
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("id")
}

func create(cmd *cobra.Command, args []string) {
	lime := chalk.Green.NewStyle().
		WithBackground(chalk.Black).
		WithTextStyle(chalk.Bold).
		Style
	fmt.Println(lime("New user has been created"))
}
