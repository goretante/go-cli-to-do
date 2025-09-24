/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-to-do",
	Short: "To do CLI Application",
	Long: `go-cli-to-do is a simple and efficient command-line tool designed to help you manage your tasks.

With this CLI application you can:
- Add new tasks
- List all your tasks
- Mark tasks as completed
- Remove tasks
- etc.

Built with Go and Cobra library.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(infoCmd)

	addCmd.Flags().StringVarP(&dueDate, "date", "d", "", "Date until the task needs to be done (format: YYYY-MM-DD)")
}
