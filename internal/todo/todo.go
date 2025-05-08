package todo

import "go-todo/internal/models"

type TodoList struct {
	todoList []models.TodoTask
}

func (t *TodoList) AddTodo(task models.TodoTask) {
	t.todoList = append(t.todoList, task)
}
