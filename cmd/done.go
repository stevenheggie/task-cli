/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stevenheggie/task/database"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task on your to-do list as complete",
	Long: `$ task done 1
"Task "wash the dishes" has been marked as complete."`,
	Run: func(cmd *cobra.Command, args []string) {
		database.MarkTodoDone("./todo.db", args[0])
		fmt.Printf("Task \"%s\" has been marked as complete.", args[0])
		database.ReorderKeys("./todo.db")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
