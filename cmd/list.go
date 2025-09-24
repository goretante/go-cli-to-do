/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
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

		now := time.Now()

		for _, t := range tasks {
			var lineColor *color.Color

			status := "✖"
			if t.IsDone {
				status = "✓"
			}

			hasDueDate := !t.DueDate.IsZero()

			switch {
			case t.IsDone && hasDueDate && t.UpdatedOn.After(t.DueDate):
				lineColor = color.New(color.FgBlack, color.BgYellow)
			case t.IsDone:
				lineColor = color.New(color.FgBlack, color.BgGreen)
			case !t.IsDone && hasDueDate && t.DueDate.Before(now):
				lineColor = color.New(color.FgWhite, color.BgRed)
			default:
				lineColor = color.New(color.FgBlack, color.BgWhite)
			}

			due := "N/A"
			if hasDueDate {
				due = t.DueDate.Format("2006-01-02")
			}

			lineColor.Printf("%d.\t%s\t[%s] Due: %s",
				t.ID,
				t.Description,
				status,
				due,
			)
			fmt.Println()
		}
	},
}
