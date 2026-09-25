package main

import (
	"errors"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	tasks = []Task{
		{ID: 1, Title: "Learn GitHub Actions", Completed: false},
		{ID: 2, Title: "Build Docker image", Completed: false},
	}

	nextID = 3
	mu     sync.Mutex
)

func getTasks() []Task {
	mu.Lock()
	defer mu.Unlock()

	return tasks
}

func createTask(title string) Task {
	mu.Lock()
	defer mu.Unlock()

	task := Task{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}

	nextID++
	tasks = append(tasks, task)

	return task
}

func findTask(id int) (*Task, error) {
	mu.Lock()
	defer mu.Unlock()

	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}

	return nil, errors.New("task not found")
}
