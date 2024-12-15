package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=service_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных:", err)
	} else {
		fmt.Println("Успешное подключение к базе данных!")
	}
}
