/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks with their status.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for i, task := range tasks {
			status := "✖"
			if task.IsDone {
				status = "✓"
			}
			fmt.Printf("%d.\t%s\t[%s]\n", i+1, task.Description, status)
		}
	},
}
