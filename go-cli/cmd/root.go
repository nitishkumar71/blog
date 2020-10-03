package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "go-cli is sample cli tool for demo purpose ",
	Long: `go-cli is an sample cli tool being built for demo purpose.
		It will be used to give demo og cobra library`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("PProvide sub-command")
	},
}

func init() {
	rootCmd.AddCommand(createNewFileCommand())
	rootCmd.AddCommand(createNewDirectoryCommand())
}

// Execute functionn is the entry point for command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
