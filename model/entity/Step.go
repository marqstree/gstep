package entity

type Step struct {
	Id          int
	Title       string
	Category    string
	Form        map[string]any
	AuditMethod string
	Expression  string
	NextSteps   []Step
	PrevStepIds []int
}

func (e Step) TableName() string {
	return "step"
}

func (e Step) GetId() any {
	return e.Id
}
