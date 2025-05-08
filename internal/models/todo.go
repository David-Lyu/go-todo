package models

type TodoTask struct {
	Task          string
	StartDateTime string
	EndDateTime   string
	Alert         bool
}

type Task struct {
	Subtask TodoTask
	TodoTask
}
