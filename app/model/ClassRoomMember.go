package model

type ClassRoomMembers struct {
	Id          int `gorm:"column:id;primary_key"`
	UserId      int `gorm:"column:user_id"`
	ClassRoomId int `gorm:"column:classroom_id"`
}

func (ClassRoomMembers) TableName() string {
	return "classroom_members"
}
