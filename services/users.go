package services

import (
	"ifp-analysis.com/database"
	"ifp-analysis.com/models"
)

// ListAllUsers retrieves all users from the database
func ListAllUsers() ([]models.User, error) {
	var users []models.User

	// Usar a instância global DB para acessar o banco
	result := database.DB.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
