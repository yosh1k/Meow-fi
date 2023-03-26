package services

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"errors"

	"gorm.io/gorm"
)

type UserController struct {
	Interactor controller.UserInteractor
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		Interactor: controller.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(login string, password string) error {
	_, err := controller.GetUserByLogin(login)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("login already exist")
	}
	user := models.User{Login: login}
	randomSalt := auth.RandSeq()
	hashedPass := auth.HashPass(password, randomSalt, config.LocalSalt)
	user.Salt = randomSalt
	user.Password = hashedPass
	controller.Interactor.Add(user)
	return nil
}
func (controller *UserController) GetAllUsers() []models.User {
	res := controller.Interactor.GetAllUsers()
	return res
}
func (controller *UserController) GetUserByLogin(login string) (models.User, error) {
	return controller.Interactor.GetUserByLogin(login)
}
func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
