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
	if err := tasks.load(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	var rootCmd = &cobra.Command{
		Use:   "task-cli",
		Short: "A simple task tracker CLI",
	}

	var addCmd = &cobra.Command{
		Use:   "add [description]",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := tasks.add(args[0]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.save(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task added successfully\n")
		},
	}

	var updateCmd = &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Update task description",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error invalid ID: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.edit(id, args[1]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.save(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task %d updated successfully\n", id)
		},
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error invalid ID: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.delete(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.save(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task %d deleted successfully\n", id)
		},
	}

	var markInProgressCmd = &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Mark task as in progress",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error invalid ID: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.markInProgress(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.save(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Task %d marked as in-progress\n", id)
		},
	}

	var markDoneCmd = &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Mark task as done",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error invalid ID: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.markDone(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			if err := tasks.save(filename); err != nil {
				fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
				os.Exit(1)
			}
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
			if err := tasks.printList(status); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(addCmd, updateCmd, deleteCmd, markInProgressCmd, markDoneCmd, listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
