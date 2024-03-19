package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"tgBotIP/internal/types"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "store.db")
	if err != nil {
		fmt.Println(err)
	}
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id BIGINT, name TEXT);")
	if err != nil {
		return
	}
	fmt.Println(result)
}

func ExistUser(userID int64) (bool, error) {
	var result int
	exec := db.QueryRow("Select count(*) from users where id = ?", userID)
	err := exec.Scan(&result)
	if err != nil {
		return true, fmt.Errorf("db troubles")
	}
	if result != 0 {
		return true, nil
	}
	return false, nil
}

func CreateUser(user types.User) error {
	_, err := db.Exec("INSERT INTO users VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}
