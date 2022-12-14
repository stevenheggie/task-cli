/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stevenheggie/task/database"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Long: `$ task list
You have the following tasks:
1. wash dishes
2. go to the gym`,
	Run: func(cmd *cobra.Command, args []string) {

		numEntries := database.GetNumEntries("./todo.db")

		if numEntries > 0 {
			fmt.Println("You have the following tasks:")
			database.ViewTodoList("./todo.db")
		} else {
			fmt.Println("Your to-do list is empty. Woohoo!")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
