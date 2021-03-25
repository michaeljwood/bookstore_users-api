package users_db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "MYSQL_USERS_USERNAME"
	mysql_users_password = "MYSQL_USERS_PASSWORD"
	mysql_users_host     = "MYSQL_USERS_HOST"
	mysql_users_schema   = "MYSQL_USERS_SCHEMA"
)

var (
	Client   *sqlx.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	datasourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		username,
		password,
		host,
		schema,
	)
	var err error
	Client, err = sqlx.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
