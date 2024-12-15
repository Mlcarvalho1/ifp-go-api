package services

import (
	"fmt"

	"ifp-analysis.com/database"
	"ifp-analysis.com/models"
)

func ListAllUsers() ([]models.User, error) {
	var users []models.User

	result := database.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func GetCurrentUser(userId *int) (models.User, error) {
	var user models.User

	if userId == nil {
		return models.User{}, fmt.Errorf("userId is nil")
	}

	result := database.DB.First(&user, *userId)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
