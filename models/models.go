package models

import (
	"log"
	"login/config"
)

//User user
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//CreateUser get new job id for old sid
func CreateUser(user User) (int, error) {
	db := config.GetDB()

	sqlInsert := `INSERT INTO account(username,email,password) VALUES($1,$2,$3) returning user_id`

	var userID int

	err := db.QueryRow(sqlInsert, user.Name, user.Email, user.Password).Scan(&userID)

	if err != nil {
		log.Println("Error in CreateUser", err.Error())
	}
	return userID, err
}

//GetUser get new job id for old sid
func GetUser(user User) (bool, User) {
	db := config.GetDB()

	// rows := db.QueryRow("SELECT user_id, name, password from account where email = " + user.Email)
	sqlInsert := `SELECT user_id, username, password from account where email = $1`

	var userID int64
	var name string
	var password string

	err := db.QueryRow(sqlInsert, user.Email).Scan(&userID, &name, &password)

	if err != nil {
		log.Println("Error in GetUser", err.Error())
	}
	var flag bool
	if password == user.Password {
		flag = true
		user.ID = userID
		user.Name = name
	}
	// flag = true correct password
	// flag = false wrong password
	return flag, user
}

// CheckUser check for user email in database
func CheckUser(user User) bool {
	db := config.GetDB()

	// rows := db.QueryRow("SELECT count(*) from account where email = $1;" + user.Email)

	var count int

	sqlInsert := `SELECT count(*) from account where email = $1;`

	err := db.QueryRow(sqlInsert, user.Email).Scan(&count)

	if err != nil {
		log.Println("Error in CheckUser", err.Error())
	}
	var flag bool
	if count > 0 {
		flag = true
	}
	// flag = true user exist
	// flag = flase user not exist
	return flag
}
