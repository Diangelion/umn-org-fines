package repositories

import (
	"backend/internal/models"
	"backend/utils"
	"database/sql"
	"fmt"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) SaveCredential(userID string, user *models.UserRegistration) {
    // Insert user into the database
    query := "INSERT INTO user_credentials (user_id, password) VALUES ($1, $2)"
    r.db.QueryRow(query, userID, user.Password)
}

func (r *UserRepository) CreateUser(user *models.UserRegistration) error {
    var userID string

    // Insert user into the database
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    if errCreate := r.db.QueryRow(query, user.Name, user.Email).Scan(&userID); errCreate != nil {
        return errCreate
    }

    r.SaveCredential(userID, user)
    return nil
}


func (r *UserRepository) CheckCredential(user *models.UserLogin) error {
    var hashedPassword string

    query := `
        SELECT password 
        FROM user_credentials 
        WHERE user_id = (SELECT id FROM users WHERE email = $1)
    `
    if errGetPassword := r.db.QueryRow(query, user.Email).Scan(&hashedPassword); errGetPassword != nil {
        if errGetPassword == sql.ErrNoRows {
            return fmt.Errorf("invalid credentials")
        }
        return fmt.Errorf("unable to process the request")
    }

    if isValidPassword := utils.CheckPasswordHash(user.Password, hashedPassword); !isValidPassword {
        return fmt.Errorf("invalid credentials")
    }

    return nil
}