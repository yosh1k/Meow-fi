package repo

import "Meow-fi/internal/models"

type UserRepository interface {
	Store(models.User)
	Select() []models.User
	SelectById(id string) models.User
	SelectByLogin(login string) (models.User, error)
	Delete(id string)
}
