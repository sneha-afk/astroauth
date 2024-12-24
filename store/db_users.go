package store

import (
	"github.com/sneha-afk/astroauth/models"
	"github.com/sneha-afk/astroauth/utils"
)

const createUserQuery = "INSERT INTO Users VALUES (?, ?, ?, ?);"
const getHashes = "SELECT Password FROM Users WHERE ? = Username OR ? = Email;"

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

func CheckUserCredentials(attempt models.UserInternal) (bool, error) {
	rows, err := DB.Query(getHashes, attempt.Username, attempt.Email)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var hash string
		err = rows.Scan(&hash)
		if err != nil {
			return false, err
		}

		matched := utils.VerifyPassword([]byte(hash), attempt.Password)
		if matched {
			return true, nil
		}
	}

	err = rows.Err()
	return false, err
}
