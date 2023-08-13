package entity

type User struct {
	Id            string `json:"id" gorm:"primarykey"`
	Name          string `json:"name"`
	PositionTitle string `json:"positionTitle"`
	PositionCode  string `json:"positionCode"`
	isLeader      bool   `json:"isLeader"`
	departmentId  int    `json:"departmentId"`
}

func (e User) TableName() string {
	return "user"
}

func (e User) GetId() any {
	return e.Id
}
