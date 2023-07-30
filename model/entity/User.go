package entity

type User struct {
	Id           string `gorm:"primarykey"`
	Name         string
	Position     string
	isLeader     bool
	departmentId int
}

func (e User) TableName() string {
	return "user"
}

func (e User) GetId() any {
	return e.Id
}
