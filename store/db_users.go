package store

import (
	"github.com/sneha-afk/astroauth/models"
)

const createUserQuery = "INSERT INTO Users VALUES (?, ?, ?, ?);"

func CreateUser(u models.UserInternal) error {
	tx, err := DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Preparex(createUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.Username, u.Email, u.Password)
	if err != nil {
		return err
	}

	// Any deferred statements after will not impact
	tx.Commit()

	return nil
}
