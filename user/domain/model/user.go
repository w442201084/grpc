package model

type User struct {
	// 主键
	Id int64 `gorm:"primary_key;not_null;auto_increment"`

	// 用户名称
	UserName string `gorm:"not_null"`

	//
	FirstName string `gorm:"not_null"`

	Pwd string `gorm:"not_null"`

	HashPwd string `gorm:"not_null"`
}
