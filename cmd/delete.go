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

var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete task.",
	Long:  "Delete selected task.",
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

		taskDescription := task.Description
		result = db.DB.Delete(&task)
		if result.Error != nil {
			fmt.Printf("Failed to delete task \"%s\": %v.\n", taskDescription, result.Error)
			return
		}
		fmt.Printf("Deleted \"%s\" task.\n", taskDescription)
	},
}
