package vo

import "github.com/marqstree/gstep/model/entity"

type CandidateVo struct {
	entity.Candidate
	Department entity.Department
	User       entity.User
	Position   string
}
