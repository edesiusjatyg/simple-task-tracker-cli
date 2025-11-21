package main

import "time"

type Task struct{
	Id int
	Description string
	Status string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Tasks []Task

func (tasks *Tasks) add(description string){
	task := Task{
		Id: len(*tasks) + 1,
		Description: description,
		Status: "not-done",
		CreatedAt: time.Now(),
	}

	*tasks = append(*tasks, task)
}

func (tasks *Tasks) edit(id int, description string){
	for i, task := range *tasks{
		if task.Id == id{
			now := time.Now()
			(*tasks)[i].Description = description
			(*tasks)[i].UpdatedAt = &now
		}
	}
}

func (tasks *Tasks) delete(id int){
	for i, task := range *tasks{
		if task.Id == id{
			temp := *tasks
			*tasks = append(temp[:i], temp[i+1:]...)
		}
	}
}

func (tasks *Tasks) markInProgress(id int){
	for i, task := range *tasks{
		if task.Id == id{
			now := time.Now()
			(*tasks)[i].Status = "in-progress"
			(*tasks)[i].UpdatedAt = &now
		}
	}
}

func (tasks *Tasks) markDone(id int){
	for i, task := range *tasks{
		if task.Id == id{
			now := time.Now()
			(*tasks)[i].Status = "done"
			(*tasks)[i].UpdatedAt = &now
		}
	}
}

func (tasks *Tasks) list(status *string) Tasks{
	if status == nil{
		return *tasks
	}

	filtered := Tasks{}
	for _, task := range *tasks{
		if task.Status == *status {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

