package dto

type TaskRefuseDto struct {
	TaskId     int
	Form       *map[string]any
	UserId     string
	PrevStepId int
}
