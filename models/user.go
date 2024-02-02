package models

import (
	"errors"
	"restwithjwt/db"
	"restwithjwt/utils"
)

type User struct {
	ID       int64
	EMAIL    string `binding:"required"`
	PASSWORD string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.PASSWORD)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.EMAIL, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.EMAIL)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		print(err.Error())
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, u.PASSWORD)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
