/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/goretante/go-cli-to-do/internal/db"
	"github.com/goretante/go-cli-to-do/internal/models"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [task number]",
	Short: "Prints task information.",
	Long:  "Prints task description, task state, task creation date & time, task state update date & time",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			fmt.Println("Invalid task ID.")
			return
		}

		var task models.Task
		result := db.DB.First(&task, id)
		if result.Error != nil {
			fmt.Printf("Task with ID %d not found.\n", id)
			return
		}

		taskStatus := "Ongoing"
		if task.IsDone {
			taskStatus = "Done"
		}
		fmt.Printf(`%d. %s
  - State: %s
  - Task created: %s
  - Task updated: %s
`,
			task.ID, task.Description,
			taskStatus,
			task.CreatedOn.Format("2006-01-02 15:04:05"),
			task.UpdatedOn.Format("2006-01-02 15:04:05"))
	},
}
