package models

import (
	"database/sql"
	"log"
	"login/config"
)

//User user
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//CreateUser get new job id for old sid
func CreateUser(user User) int {
	db := config.GetDB()

	rows := db.QueryRow("INSERT INTO db_user(name,email,password) "+user.Name, user.Email, user.Password)

	var nid int
	err := rows.Scan(&nid)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			// no such job id
			return 0
		}
	}
	return nid
}

//GetUser get new job id for old sid
func GetUser(user User) int {
	db := config.GetDB()

	rows := db.QueryRow("SELECT id, password from db_user where email = " + user.Email)

	var nid int
	err := rows.Scan(&nid)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			// no such job id
			return 0
		}
	}
	return nid
}
