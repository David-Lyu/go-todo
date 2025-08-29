package todo

import "go-todo/internal/models"

type Todo struct {
	TodoList models.TodoList
}

func (t *Todo) AddTodo(task models.Task) {
	t.TodoList = append(t.TodoList, task)
}
