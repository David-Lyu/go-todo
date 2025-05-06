package config

import "go-todo/internal/models"

type Task struct {
	subtask []models.TodoTask
	models.TodoTask
}

/*
Checks to see if json config exists, if not it creates it
*/
func (t *Task) Init() {

}
