package database

import (
	"database/sql"
	"fmt"
	"tgBotIP/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

type Databaser interface {
	ExistUser(int64) (bool, error)
	CreateUser(types.User) error
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./database/store.db") //
	if err != nil {
		panic("can not open db")
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	result, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id BIGINT, name TEXT);")
	if err != nil {
		return
	}
	fmt.Println(result)
}

func ExistUser(userID int64) (bool, error) {
	var result int
	exec := db.QueryRow("SELECT COUNT(*) FROM users WHERE id = ?", userID)
	err := exec.Scan(&result)
	if err != nil {
		fmt.Println("db troubles:", err)
		return false, err // Возвращаем false и ошибку
	}
	if result != 0 {
		return true, nil // Возвращаем true и nil (без ошибки)
	}
	return false, nil // Возвращаем false и nil (без ошибки)
}
func CreateUser(user types.User) error {
	_, err := db.Exec("INSERT INTO users VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		return err
	}
	return nil
}
