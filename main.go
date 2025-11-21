package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/spf13/cobra"
)

const filename = "tasks.json"

var tasks Tasks

func main() {
	tasks.load(filename)

	var rootCmd = &cobra.Command{
		Use:   "task-cli",
		Short: "A simple task tracker CLI",
	}

	var addCmd = &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			tasks.add(args[0])
			tasks.save(filename)
			fmt.Printf("Task added successfully (ID: %d)\n", len(tasks))
		},
	}

	var updateCmd = &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Update task description",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			tasks.edit(id, args[1])
			tasks.save(filename)
			fmt.Printf("Task %d updated successfully\n", id)
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			tasks.delete(id)
			tasks.save(filename)
			fmt.Printf("Task %d deleted successfully\n", id)
		},
	}

	var markInProgressCmd = &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Mark task as in progress",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			tasks.markInProgress(id)
			tasks.save(filename)
			fmt.Printf("Task %d marked as in-progress\n", id)
		},
	}

	var markDoneCmd = &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Mark task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			tasks.markDone(id)
			tasks.save(filename)
			fmt.Printf("Task %d marked as done\n", id)
		},
	}

	var listCmd = &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks or filter by status (done, in-progress, not-done)",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var status *string
			if len(args) > 0 {
				status = &args[0]
			}
			tasks.printList(status)
		},
	}

	rootCmd.AddCommand(addCmd, updateCmd, deleteCmd, markInProgressCmd, markDoneCmd, listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
