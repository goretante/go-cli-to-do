/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/goretante/go-cli-to-do/internal/db"
	"github.com/goretante/go-cli-to-do/internal/models"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Long:  "Add a new task with a description to your task list.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		task := models.Task{Description: description}

		if err := db.DB.Create(&task).Error; err != nil {
			fmt.Printf("Failed to add task: %v\n", err)
			return
		}
		fmt.Println("Task added:", description)
	},
}
