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
		Status: "in-progress",
		CreatedAt: time.Now(),
	}

	*tasks = append(*tasks, task)
}

func (tasks *Tasks) edit(id int, description string){
	for i, task := range *tasks{
		if task.Id == id{
			(*tasks)[i].Description = description
		}
	}
}



