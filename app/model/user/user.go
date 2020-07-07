package user

type User struct {
	Id int `gorm:"column:id;primary_key"`
	Name string  `gorm:"column:name"`
}