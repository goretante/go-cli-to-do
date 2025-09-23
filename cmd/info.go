/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [task number]",
	Short: "Prints task information.",
	Long:  "Prints task description, task state, task creation date & time, task state update date & time",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if index < 1 || err != nil {
			fmt.Println("Invalid task number entered.")
			return
		}

		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("Task number out of tasks range.")
			return
		}

		taskStatus := "Ongoing"
		if tasks[index-1].IsDone {
			taskStatus = "Done"
		}
		fmt.Printf(`%d. %s
  - State: %s
  - Task created: %s
  - Task updated: %s
`,
			index, tasks[index-1].Description,
			taskStatus,
			tasks[index-1].CreatedOn.Format("2006-01-02 15:04:05"),
			tasks[index-1].UpdatedOn.Format("2006-01-02 15:04:05"))
	},
}
