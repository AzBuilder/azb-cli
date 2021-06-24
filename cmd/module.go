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

	"github.com/spf13/cobra"
)

var moduleLong = `
This command consists of multiple subcommands to interact with modules.
It can be used to create, update, delete and list modules.
`

var moduleCmd = &cobra.Command{
	Use:   "module create|update|delete|list [ARGS]",
	Short: "create, update, delete and list modules",
	Long:  moduleLong,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("module called")
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)
}
