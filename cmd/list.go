/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"

	"github.com/goretante/go-cli-to-do/internal/db"
	"github.com/goretante/go-cli-to-do/internal/models"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "List all tasks with their status.",
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []models.Task
		if err := db.DB.Find(&tasks).Error; err != nil {
			fmt.Printf("Failed to list tasks: %v", err)
			return
		}

		for _, t := range tasks {
			status := "✖"
			if t.IsDone {
				status = "✓"
			}
			fmt.Printf("%d.\t%s\t[%s]\n", t.ID, t.Description, status)
		}
	},
}
