package repo

import "Meow-fi/internal/models"

type UserRepository interface {
	Store(models.User)
	Select() []models.User
	Delete(id string)
}
