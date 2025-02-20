package repositories

import (
	"backend/internal/models"
	"backend/utils"
	"database/sql"
	"errors"
	"log"

	"github.com/lib/pq"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) SaveCredential(userID string, user *models.UserRegistration) error {
    // Insert user into the database
    query := "INSERT INTO user_credentials (user_id, password) VALUES ($1, $2)"
    if _, err := r.db.Exec(query, userID, user.Password); err != nil {
        log.Println("SaveCredential | Insert user credentials error: ", err)
        return errors.New("Unable to create credentials.")
    }
    return nil
}

func (r *UserRepository) CreateUser(user *models.UserRegistration) error {
    var userId string

    // Insert user into the database
    query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
    if err := r.db.QueryRow(query, user.Name, user.Email).Scan(&userId); err != nil {
        // Check if the error is from a unique constraint
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "23505" { // Unique violation error code
				return &models.DuplicateEmailError{Email: user.Email}
			}
		}
        log.Println("CreateUser | Insert users error: ", err)
        return errors.New("Unable to create user.")
    }
    return r.SaveCredential(userId, user)
}


func (r *UserRepository) CheckCredential(user *models.UserLogin) (string, error) {
    var userId string
    var hashedPassword string

    query := `
        SELECT users.id, user_credentials.password 
	    FROM users JOIN user_credentials 
        ON users.id = user_credentials.user_id
		WHERE users.email = $1
    `
    if err := r.db.QueryRow(query, user.Email).Scan(&userId, &hashedPassword); err != nil {
        if err == sql.ErrNoRows {
            return "", errors.New("Invalid email and/or password.")
        }
        log.Println(err)
        return "", errors.New("unable to process the request")
    }

    if isValidPassword := utils.CheckPasswordHash(user.Password, hashedPassword); !isValidPassword {
        return "", errors.New("Invalid email and/or password.")
    }

    return userId, nil
}