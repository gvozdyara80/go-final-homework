package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-final-homework/models"
)

func (r *Repository) CreateUser(user models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (username, first_name, last_name, email, password) VALUES (?, ?, ?, ?, ?) RETURNING id",
		user.Username, 
		user.Firstname, 
		user.Lastname, 
		user.Email, 
		user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsInUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (r *Repository) GetUserByID(id int) (*models.User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(models.User)
	for rows.Next() {
		u, err = scanRowsInUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.Id == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func scanRowsInUser(rows *sql.Rows) (*models.User, error) {
	user := new(models.User)

	err := rows.Scan(
		&user.Id,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}