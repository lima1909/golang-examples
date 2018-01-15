package cmd

import (
	"github.com/spf13/cobra"
)

const defaultBashCompletionFileName = "bash-completion.sh"

// bashCompletionCmd represents the bashCompletion command
var bashCompletionCmd = &cobra.Command{
	Use:   "bash-completion [file name]",
	Short: "you can generate a file for bash completion",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		createBashCompletionFile(getBashCompletionFileByArgs(args))
	},
}

func init() {
	RootCmd.AddCommand(bashCompletionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bashCompletionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bashCompletionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getBashCompletionFileByArgs(args []string) string {
	if args == nil || len(args) == 0 {
		return defaultBashCompletionFileName
	}
	return args[0]
}

func createBashCompletionFile(fileName string) error {
	err := RootCmd.GenBashCompletionFile(fileName)
	return err
}
