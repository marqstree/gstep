package entity

type Step struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Category    string         `json:"category"`
	Form        map[string]any `json:"form"`
	Candidates  []Candidate    `json:"candidates"`
	Expression  string         `json:"expression"`
	AuditMethod string         `json:"auditMethod"`
	BranchSteps []*Step        `json:"branchSteps"`
	NextStep    *Step          `json:"nextStep"`
}
