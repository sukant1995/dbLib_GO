package dblib

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//Local DB
	dbUser     = "root"
	dbPassword = "root"
	dbName     = "mydb"
	dbHost     = "localhost"
	dbPort     = "3306"
)

var db *sql.DB
var minutes int

func MySqlInit() {
	vdb, err := sql.Open("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("I am able to connect to Database")
	}
	//vdb.SetConnMaxLifetime(-1);
	vdb.SetConnMaxLifetime(time.Second)
	vdb.SetMaxIdleConns(0)
	vdb.SetMaxOpenConns(5)
	//defer db.Close()
	// make sure connection is available
	err = vdb.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	db = vdb
	minutes = 0
	time.Sleep(50 * time.Millisecond)
}

func SetConfig(user, pass, dbname, host, port string) {
	dbUser = user
	dbPassword = pass
	dbName = dbname
	dbHost = host
	dbPort = port
}
