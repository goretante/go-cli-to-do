/*
Copyright © 2025 Ante Goreta agoret00@fesb.hr
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete task.",
	Long:  "Delete selected task.",
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

		taskDescription := tasks[index-1].Description
		tasks = append(tasks[:index-1], tasks[index:]...)
		saveTasks(tasks)
		fmt.Printf("Deleted \"%s\" task.\n", taskDescription)
	},
}
