/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark task as done.",
	Long:  "Mark selected task as done.",
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

		tasks[index-1].IsDone = true
		tasks[index-1].UpdatedOn = time.Now()
		saveTasks(tasks)
		fmt.Printf("Task %d. marked as done.\n", index)
	},
}
