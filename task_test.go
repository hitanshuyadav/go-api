package main

import "testing"

func TestCreateTask(t *testing.T) {

	task := createTask("Learn Docker")

	if task.Title != "Learn Docker" {
		t.Errorf("expected title 'Learn Docker', got '%s'", task.Title)
	}

	if task.ID == 0 {
		t.Error("expected task ID to be generated")
	}
}

func TestFindTask(t *testing.T) {

	task, err := findTask(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
}

func TestFindTaskNotFound(t *testing.T) {

	_, err := findTask(9999)

	if err == nil {
		t.Error("expected error for missing task")
	}
}
