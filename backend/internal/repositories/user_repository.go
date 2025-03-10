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
                log.Println("CreateUser | Failed to insert user")
				return &models.DuplicateEmailError{Email: user.Email}
			}
		}
        log.Println("CreateUser | Insert user error: ", err)
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
            log.Println("CheckCredential | User not found")
            return "", errors.New("Invalid email and/or password.")
        }
        log.Println("CheckCredential | Get user id and password error: ", err)
        return "", errors.New("unable to process the request")
    }

    if isValidPassword := utils.CheckPasswordHash(user.Password, hashedPassword); !isValidPassword {
        log.Println("CheckCredential | Password and hashed password don't match")
        return "", errors.New("Invalid email and/or password.")
    }

    return userId, nil
}

func (r *UserRepository) GetUser(userId string) (*models.UserEdit, error) {
    query := `
        SELECT name, email, profile_photo, cover_photo
        FROM users
        WHERE id = $1
    `

    var user models.UserEdit
    if err := r.db.QueryRow(query, userId).Scan(&user.Name, &user.Email, &user.ProfilePhoto, &user.CoverPhoto); err != nil {
        if err == sql.ErrNoRows {
            // Handle the case where no rows were found
            log.Printf("GetUser | Get user error: %v, with id: %s\n", err, userId)
            return nil, errors.New("Unable to get user (not found).")
        }
        log.Println("GetUser | Error getting user: ", err)
        return nil, errors.New("Unable to get user.")
    }

    return &user, nil
}


func (r *UserRepository) UpdateUser(user *models.UserEdit, userId string) error {
    query := `
        UPDATE users
        SET name = $1, email = $2, profile_photo = $3, cover_photo = $4
        WHERE id = $5
    `

    // Use Exec for UPDATE queries
    result, err := r.db.Exec(query, user.Name, user.Email, user.ProfilePhoto, user.CoverPhoto, userId)
    if err != nil {
        log.Println("Error updating user:", err)
        return errors.New("UpdateUser | Unable to update user.")
    }

    // Optional: Check how many rows were affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Println("Error getting rows affected:", err)
        return errors.New("UpdateUser | Unable to get rows affected.")
    }

    if rowsAffected == 0 {
        log.Println("No rows affected; user ID may not exist")
        return errors.New("UpdateUser | No user found (0 row affected).")
    }

    return nil
}