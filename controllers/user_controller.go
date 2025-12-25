package controllers

import (
	"github.com/amaricee/simple_crud/config"
	"github.com/amaricee/simple_crud/models"
)

func CreateUser(user models.User) error {
	query := `
		INSERT INTO users (name, email)
		VALUES (?, ?)
	`

	_, err := config.DB.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
	`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(id int, user models.User) (bool, error) {
	query := `
		UPDATE users
		SET name = ?, email = ?
		WHERE id = ?
	`

	result, err := config.DB.Exec(query, user.Name, user.Email, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil // user tidak ditemukan
	}

	return true, nil
}

func DeleteUser(id int) (bool, error) {
	query := `
		DELETE FROM users
		WHERE id = ?
	`

	result, err := config.DB.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil // user tidak ditemukan
	}

	return true, nil
}
