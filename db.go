package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "store.db")
	if err != nil {
		fmt.Println(err)
	}
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id BIGINT, name TEXT, password TEXT);")
	if err != nil {
		return
	}
	fmt.Println(result)
}

func createUser(userID, name, password string) (bool, error) {
	exist, _ := existUser(userID)
	if !exist {
		_, err := db.Exec("INSERT INTO users VALUES (?, ?, ?)", userID, name, password)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, fmt.Errorf("user already exists")
}

func existUser(userID string) (bool, error) {
	var result int
	exec := db.QueryRow("Select count(*) from users where id = ?", userID)
	err := exec.Scan(&result)
	if err != nil {
		return true, fmt.Errorf("user already exists")
	}
	if result != 0 {
		return true, fmt.Errorf("user already exists")
	}
	return false, nil
}
