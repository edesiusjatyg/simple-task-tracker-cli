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

func (tasks *Tasks) add(description string) error {
	if description == "" {
		return fmt.Errorf("task description cannot be empty")
	}
	
	maxId := 0
	for _, task := range *tasks{
		if task.Id > maxId{
			maxId = task.Id
		}
	}

	task := Task{
		Id:          maxId + 1,
		Description: description,
		Status:      "not-done",
		CreatedAt:   time.Now(),
	}

	*tasks = append(*tasks, task)
	return nil
}

func (tasks *Tasks) edit(id int, description string) error {
	if id <= 0{
		return fmt.Errorf("invalid ID")
	}
	if description == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	for i, task := range *tasks {
		if task.Id == id {
			now := time.Now()
			(*tasks)[i].Description = description
			(*tasks)[i].UpdatedAt = &now
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tasks *Tasks) delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	for i, task := range *tasks {
		if task.Id == id {
			temp := *tasks
			*tasks = append(temp[:i], temp[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tasks *Tasks) markInProgress(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	for i, task := range *tasks {
		if task.Id == id {
			now := time.Now()
			(*tasks)[i].Status = "in-progress"
			(*tasks)[i].UpdatedAt = &now
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tasks *Tasks) markDone(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid ID")
	}

	for i, task := range *tasks {
		if task.Id == id {
			now := time.Now()
			(*tasks)[i].Status = "done"
			(*tasks)[i].UpdatedAt = &now
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (tasks *Tasks) list(status *string) (Tasks, error) {
	statuses := map[string]bool{
		"": true,
		"done": true,
		"in-progress": true,
		"not-done": true,
	}

	if status == nil {
		return *tasks, nil
	}

	if !statuses[*status]{
		return nil, fmt.Errorf("invalid status '%s' statuses: done, in-progress, not-done", *status)
	}

	filtered := Tasks{}
	for _, task := range *tasks {
		if task.Status == *status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (tasks *Tasks) save(filename string) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func (tasks *Tasks) load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err){
			*tasks = Tasks{}
			return nil
		}
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		*tasks = Tasks{}
		return nil
	}

	if err := json.Unmarshal(data, tasks); err != nil {
		return fmt.Errorf("failed to unmarshal tasks: %w", err)
	}
	
	return nil
}

func (tasks *Tasks) printList(status *string) error {	
	filtered, err := tasks.list(status)
	if err != nil {
		return err
	}
	if len(filtered) == 0 {
		fmt.Println("No tasks found.")
		return nil
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

	return nil
}
