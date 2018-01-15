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

	// if a typo error occur, like this, than can guess this command
	SuggestFor: []string{"bas-completi"},
	// a short cut for this command
	Aliases: []string{"bc"},
}

func init() {
	RootCmd.AddCommand(bashCompletionCmd)
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
