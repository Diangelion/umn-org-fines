package repositories

import (
	"backend/internal/models"
	"database/sql"
	"errors"
	"log"
)

type OrganizationRepository struct {
    db *sql.DB
}

func NewOrganizationRepository(db *sql.DB) *OrganizationRepository {
    return &OrganizationRepository{db}
}

func (r *OrganizationRepository) GetListOrganizationRepository(orgList *models.GetListOrganization, userId string) error {
    query := `
        SELECT name, photo, descriptions, start_date, end_date
        FROM organizations org
        JOIN user_organizations usr_orgs ON org.id = usr_orgs.organization_id
        WHERE usr_orgs.user_id = $1
    `

    // Execute the query
    rows, err := r.db.Query(query, userId)
    if err != nil {
        log.Println("GetListOrganizationRepository | Error querying database:", err)
        return errors.New("Unable to get list organization.")
    }
    defer rows.Close() // Ensure rows are closed after iteration

    // Always initialize as an empty slice to avoid nil
    orgList.List = []models.CreateOrganization{}

    // Iterate over the rows and append each organization name to the slice
    for rows.Next() {
        var org models.CreateOrganization
        if err := rows.Scan(
                    &org.OrganizationName,
                    &org.OrganizationPhoto,
                    &org.OrganizationDescriptions,
                    &org.OrganizationStartDate,
                    &org.OrganizationEndDate,
        ); err != nil {
            log.Println("GetListOrganizationRepository | Error scanning row:", err)
            return errors.New("Unable to process list organization.")
        }
        orgList.List = append(orgList.List, org)
    }

    // Check for errors after iteration
    if err := rows.Err(); err != nil {
        log.Println("GetListOrganizationRepository | Error after iterating rows:", err)
        return errors.New("Unable to get list organization.")
    }

    return nil
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
    // query := "INSERT INTO organizations (name, descriptions, organization_photo) VALUES ($1, $2, $3) RETURNING id"
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
    return nil
}