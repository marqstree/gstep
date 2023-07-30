package dto

type TaskDto struct {
	ProcessId int
	Form      *map[string]any `json:"form,omitempty"`
	Assignees *[]string
	StepId    int
	State     string
}
