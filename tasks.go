package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type Tasks []Task

func (tasks *Tasks) add(description string) {
	task := Task{
		Id:          len(*tasks) + 1,
		Description: description,
		Status:      "not-done",
		CreatedAt:   time.Now(),
	}

	*tasks = append(*tasks, task)
}

func (tasks *Tasks) edit(id int, description string) {
	for i, task := range *tasks {
		if task.Id == id {
			now := time.Now()
			(*tasks)[i].Description = description
			(*tasks)[i].UpdatedAt = &now
		}
	}
}

func (tasks *Tasks) delete(id int) {
	for i, task := range *tasks {
		if task.Id == id {
			temp := *tasks
			*tasks = append(temp[:i], temp[i+1:]...)
		}
	}
}

func (tasks *Tasks) markInProgress(id int) {
	for i, task := range *tasks {
		if task.Id == id {
			now := time.Now()
			(*tasks)[i].Status = "in-progress"
			(*tasks)[i].UpdatedAt = &now
		}
	}
}

func (tasks *Tasks) markDone(id int) {
	for i, task := range *tasks {
		if task.Id == id {
			now := time.Now()
			(*tasks)[i].Status = "done"
			(*tasks)[i].UpdatedAt = &now
		}
	}
}

func (tasks *Tasks) list(status *string) Tasks {
	if status == nil {
		return *tasks
	}

	filtered := Tasks{}
	for _, task := range *tasks {
		if task.Status == *status {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func (tasks *Tasks) save(filename string) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(filename, data, 0644)
}

func (tasks *Tasks) load(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	json.Unmarshal(data, tasks)
}

func (tasks *Tasks) printList(status *string) {
	filtered := tasks.list(status)
	if len(filtered) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Printf("%-5s %-20s %-15s %-25s %-25s\n", "ID", "Description", "Status", "Created At", "Updated At")
	fmt.Println("----------------------------------------------------------------------------------------------------")
	for _, task := range filtered {
		updatedAt := "N/A"
		if task.UpdatedAt != nil {
			updatedAt = task.UpdatedAt.Format("2006-01-02 15:04:05")
		}
		fmt.Printf("%-5d %-20s %-15s %-25s %-25s\n",
			task.Id,
			task.Description,
			task.Status,
			task.CreatedAt.Format("2006-01-02 15:04:05"),
			updatedAt)
	}
}
