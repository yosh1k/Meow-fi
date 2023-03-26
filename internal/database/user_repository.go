package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	interfaces.SqlHandler
}

func (db *UserRepository) Store(u models.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []models.User {
	var user []models.User
	db.FindAll(&user)
	return user
}

func (db *UserRepository) SelectByLogin(login string) (models.User, error) {
	var user models.User
	res := db.Where("login = ?", login).Find(&user)
	if res.Error != nil {
		return user, res.Error
	}
	if res.RowsAffected == 0 {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (db *UserRepository) SelectById(id string) models.User {
	var user models.User
	db.Where("id = ?", id).Find(&user)
	return user
}

func (db *UserRepository) Delete(id string) {
	var user []models.User
	db.DeleteById(&user, id)
}
