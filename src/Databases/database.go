package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_net_http")

	if err != nil {
		panic(err.Error())
	}else{
		fmt.Println("Database connected")
	}

	defer db.Close()
}