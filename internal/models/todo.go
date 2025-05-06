package models

type TodoTask struct {
	task          string
	startDateTime string
	endDateTime   string
	alert         bool
}

// type TodoParentTask struct {
// 	subtask []TodoTask
// 	TodoTask
// }
