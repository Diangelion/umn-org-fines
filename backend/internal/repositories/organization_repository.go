package repositories

import (
	"backend/internal/models"
	"database/sql"
)

type OrganizationRepository struct {
    db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *OrganizationRepository {
    return &OrganizationRepository{db}
}

// func (r *UserRepository) SaveCredential(userID string, user *models.RegisterUser) error {
//     // Insert user into the database
//     query := "INSERT INTO user_credentials (user_id, password) VALUES ($1, $2)"
//     if _, err := r.db.Exec(query, userID, user.Password); err != nil {
//         log.Println("SaveCredential | Insert user credentials error: ", err)
//         return errors.New("Unable to create credentials.")
//     }
//     return nil
// }

func (r *OrganizationRepository) CreateOrganizationRepository(org *models.CreateOrganization, userId string) error {
    // Insert org into the database
    query := "INSERT INTO organizations (name, descriptions, organization_photo) VALUES ($1, $2, $3) RETURNING id"
    // if err := r.db.QueryRow(query, user.Name, user.Email).Scan(&userId); err != nil {
    //     // Check if the error is from a unique constraint
	// 	if pgErr, ok := err.(*pq.Error); ok {
	// 		if pgErr.Code == "23505" { // Unique violation error code
    //             log.Println("RegisterUserRepository | Failed to insert user")
	// 			return &models.DuplicateEmailError{Email: user.Email}
	// 		}
	// 	}
    //     log.Println("RegisterUserRepository | Insert user error: ", err)
    //     return errors.New("Unable to create user.")
    // }
    // return r.SaveCredential(userId, user)
}