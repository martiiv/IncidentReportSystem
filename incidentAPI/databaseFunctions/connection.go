package databasefunctions

import (
	"database/sql"
	_ "database/sql"
	"encoding/json"
	_ "errors"
	"fmt"
	_ "fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Configuration struct {
	DB_NAME     string
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
}

// in case of further help is need https://github.com/golangbot/mysqltutorial/blob/master/select/main.go
func EstablishConnection() {

	var configuration = Configuration{}

	file, err := os.Open("config/config.development.json")
	if err != nil {
		fmt.Print(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Print(err)
	}

	cfg := mysql.Config{
		User:   configuration.DB_USERNAME,
		Passwd: configuration.DB_PASSWORD,
		Net:    "tcp",
		Addr:   configuration.DB_HOST,
		DBName: configuration.DB_NAME,
	}

	fmt.Println("Now connecting...")
	db, err := sql.Open("mydriver", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// for later testing insrt(db)
	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("connected...")
}

func insrt(dbname *sql.DB, tblname string, params []string) { //if more values are needed adjust accordingly

	stmt, err := dbname.Prepare("insert into" + tblname + " set Name=?, Lname=?,	Country=?")
	if err != nil {
		fmt.Print("helper_methods.go : 118")
		fmt.Println(err)

	}
	//just for proof of concept

	switch len(params) {
	case 1:
		_, queryError := stmt.Exec(tblname, params[1])
		if queryError != nil {
			fmt.Print("helper_methods.go : 125")
			fmt.Println(queryError)

		}

	}

	//dbname, err := con.Exec("insert into tbl (id, mdpr, isok) values (?, ?, 1)", id, mdpr)
}
