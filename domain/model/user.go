package model

type User struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`		//主键
	UserName string `gorm:"unique_index;not_null"`				//用户名称
	FirstName string
	HashPassword string
}
