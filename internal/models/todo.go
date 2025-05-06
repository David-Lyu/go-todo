package models

type TodoTask struct {
	task          string
	startDateTime string
	endDateTime   string
	alert         bool
}

type Task struct {
	subtask TodoTask
	TodoTask
}
