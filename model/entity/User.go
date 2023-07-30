package entity

type User struct {
	Id           string `gorm:"primarykey"`
	Name         string
	Position     string
	isLeader     bool
	departmentId int
}
