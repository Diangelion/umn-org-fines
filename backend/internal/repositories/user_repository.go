package repositories

import (
	"backend/internal/models"
	"database/sql"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) SaveCredential(userID string, user *models.User) {
    // Insert user into the database
    query := "INSERT INTO user_credentials (user_id, password) VALUES ($1, $2)"
    r.db.QueryRow(query, userID, user.Password)
}

func (r *UserRepository) CreateUser(user *models.User) error {
    var userID string

    // Insert user into the database
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    if errCreate := r.db.QueryRow(query, user.Name, user.Email).Scan(&userID); errCreate != nil {
        return errCreate
    }

    r.SaveCredential(userID, user)
    return nil
}