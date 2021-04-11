package database

import (
	"database/sql"
	"errors"
	"fmt"
	"user/model"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlCreateUserprofile = `
    CREATE TABLE IF NOT EXISTS userprofile(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username VARCHAR NOT NULL,
		dob VARCHAR NOT NULL,
		age VARCHAR NOT NULL,
		email VARCHAR NOT NULL,
		phonenumber VARCHAR NOT NULL
    );
	
	`
	sqlInsertUserprofile = `
	INSERT INTO userprofile 
		(username,dob,age,email,phonenumber) VALUES ('testdata','29/03/2021','23','testdata@example.com','1234567890');
    `
	sqlGetuserdetails = `
	SELECT * FROM userprofile
		WHERE username = ?
	`
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("no database found")
	}
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	err = insertTest(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func migrate(db *sql.DB) error {
	_, err := db.Exec(sqlCreateUserprofile)
	if err != nil {
		return err
	}
	return nil
}
func insertTest(db *sql.DB) error {
	_, err := db.Exec(sqlInsertUserprofile)
	if err != nil {
		return err
	}
	return nil
}

func GetUserDetails(userName string, db *sql.DB) (*model.Userprofile, error) {
	user := model.Userprofile{}
	err := db.QueryRow(sqlGetuserdetails, userName).Scan(&user.ID, &user.UserName, &user.Dob, &user.Age, &user.Email, &user.PhoneNumber)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
