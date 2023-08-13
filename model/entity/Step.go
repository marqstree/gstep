package entity

type Step struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Category    string         `json:"category"`
	Level       int            `json:"level"`
	Form        map[string]any `json:"form"`
	Candidates  []Candidate    `json:"candidates"`
	Expression  string         `json:"expression"`
	AuditMethod string         `json:"auditMethod"`
	BranchSteps []*Step        `json:"branchSteps"`
	NextStep    *Step          `json:"nextStep"`
	PrevStepIds []int          `json:"prevStepIds"`
}
