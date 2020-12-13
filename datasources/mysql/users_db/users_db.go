package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_db_username = "mysql_users_db_username"
	mysql_users_db_password = "mysql_users_db_password"
	mysql_users_db_host     = "mysql_users_db_host"
	mysql_users_db_schema   = "mysql_users_db_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_db_username)
	password = os.Getenv(mysql_users_db_password)
	host     = os.Getenv(mysql_users_db_host)
	schemaDb = os.Getenv(mysql_users_db_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schemaDb,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured")
}
