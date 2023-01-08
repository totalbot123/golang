package services

import models "users/models"

type UsersInterface interface {
	GetUsers(x, y, radius float32) []models.Users
	UpdateUsers(userName string, x, y float32) models.Users
	CreateUsers(userName string, x, y float32) models.Users
}
