/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/goretante/go-cli-to-do/internal/db"
	"github.com/goretante/go-cli-to-do/internal/models"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark task as done.",
	Long:  "Mark selected task as done.",
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

		task.IsDone = true
		task.UpdatedOn = time.Now()

		if err = db.DB.Save(&task).Error; err != nil {
			fmt.Println("Failed to mark task as done:", err)
			return
		}

		fmt.Printf("Task %d marked as done.\n", id)
	},
}
