package cmd

import (
	"fmt"

	"github.com/lima1909/golang-examples/cli/list"
	"github.com/spf13/cobra"
)

// var flagType = []string{"file", "dir"}
var flagMaxFiles int

var listCmd = &cobra.Command{
	Use:   "list [directory]",
	Short: "command to list files or directories",
	Long:  "you can get an overview from files or directories ...",

	PreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("before")
	},

	Run: func(cmd *cobra.Command, args []string) {
		list.Dir(getStartDirectoryFromArgs(args), flagMaxFiles)
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		// fmt.Println("after")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().IntVarP(&flagMaxFiles, "max-files", "m", 100, "max files list")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

}

// from wich directory show the file/directory-list
func getStartDirectoryFromArgs(args []string) string {
	fmt.Println("ARGS: ", args)
	if args != nil && len(args) > 0 {
		return args[0]
	}
	return "./"
}
