/*
Copyright © 2022 Steven Heggie github.com/stevenheggie
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI based to-do list.",
	Long:  ``,
	// 	`A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("PLPLPLPLP")
	fmt.Println(viper.GetString("DB_DIR") + "/" + viper.GetString("DB_NAME"))
}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.task-cli.yaml)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.task-cli.yaml)")

	// TODO: Add DB PATH Config option

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		fmt.Println("hello")
	} else {
		// // Find home directory.
		// home, err := os.UserHomeDir()
		cwdPath, err := os.Getwd()
		cobra.CheckErr(err)

		fmt.Println(cwdPath)
		// Search config in home directory with name ".viper-test" (without extension).
		// viper.AddConfigPath(home)

		viper.AddConfigPath(cwdPath) // use current working dir for config file (testing)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".task-cli")
		viper.SetDefault("DB_DIR", ".")        // default database directory to working dir
		viper.SetDefault("DB_NAME", "todo.db") // default database name to 'todo.db'
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
