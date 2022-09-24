/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task on your TODO list as complete",
	Long: `$ task do 1
You have completed the "wash dishes" task.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a fake \"done\" command")
		fmt.Printf("Task \"%s\" has been marked as complete.", "take out rubbish")
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
