package repositories

import (
	"database/sql"
	"umn-org-fines/internal/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    // Insert user into the database
    query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id"
    return r.db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
}
