package models

type User struct {
	UserId int    `json:"user_id" gorm:"primary_key"`
	FIO    string `json:"fio"`
}
