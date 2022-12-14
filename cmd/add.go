/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stevenheggie/task/database"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your to-do list",
	Long: `$ task add wash dishes
Added "wash dishes" to your task list.
	
$ task add go to the gym
Added "go to the gym" to your task list.`,
	Run: func(cmd *cobra.Command, args []string) {

		// create todo entry string and write to db
		entry := strings.Join(args, " ")
		database.CreateTodoEntry("./todo.db", entry)
		fmt.Printf("Added \"%s\" to your task list.", entry)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
