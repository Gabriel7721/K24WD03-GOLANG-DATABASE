package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var MySQLDB *sql.DB

func MySQLConnect() {
	// "root:root@tcp(127.0.0.1:3306)/server02_golang_db?parseTime=true"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)

	db, _ := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	MySQLDB = db
	fmt.Println("MySQL connected successfully")
}
