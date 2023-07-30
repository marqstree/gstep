package entity

type Step struct {
	Id            int
	Title         string
	Category      string
	Form          map[string]any
	Expression    string
	AgreeSteps    []Step
	RefuseStepIds []int
}
