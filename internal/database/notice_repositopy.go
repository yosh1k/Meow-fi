package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
)

type NoticeRepository struct {
	interfaces.SqlHandler
}

func (db *NoticeRepository) Store(notice models.Notice) {
	db.Create(&notice)
}
func (db *NoticeRepository) Select() []models.Notice {
	var notices []models.Notice
	db.FindAll(&notices)
	return notices
}
func (db *NoticeRepository) UpdateNotice(notice models.Notice) {
	db.Update(notice)
}
func (db *NoticeRepository) SelectById(id string) models.Notice {
	var task models.Notice
	db.Where("id = ?", id).Find(&task)
	return task
}
func (db *NoticeRepository) GetNoticeInfo(id string) models.Notice {
	var task models.Notice
	db.Preload("Client").Where("id = ?", id).Find(&task)
	return task
}
func (db *NoticeRepository) Delete(id string) {
	var notices []models.Notice
	db.DeleteById(&notices, id)
}
