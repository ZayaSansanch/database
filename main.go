package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int    `db:"id"`
	Login    string `db:"login"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Age      int    `db:"age"`
}

func main() {
	db, err := sqlx.Open("sqlite3", "data.sql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var (
		login    = gofakeit.Username()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, false, gofakeit.Number(8, 16))
		age      = gofakeit.Number(0, 100)
	)

	_, err = db.Exec("INSERT INTO users (login, email, password, age) VALUES ($1, $2, $3, $4)", login, email, password, age)
	if err != nil {
		panic(err)
	}

	rows, err := db.Queryx("SELECT id, login, email, password, age FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var u User
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u.ID, u.Email, u.Login, u.Password, u.Age)
	}
}
