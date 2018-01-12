package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd this command is called, if no command is set
var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "this my first cli",
	Long:  "this is a implementation with the spf13/cabra framework for a example of a cli solution",
	Run: func(cmd *cobra.Command, args []string) {
		// log.Printf("Command %v, with Args: %v", cmd, args)
	},
}

func init() {
}

// Execute this is the entry point from the cli.main Method ("the new main")
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
