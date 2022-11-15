package databasefunctions

import (
	"context"
	"database/sql"
	_ "errors"
	"fmt"
	apitools "incidentAPI/apiTools"
	"incidentAPI/config"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
)

/*
* File connection.go, contains basic database functionality for the software
* Also contains some database functions (Creating a system manager)
? Last revision Martin Iversen 15.11.2022
*/

// Struct Configuration used for getting environment variables for database
type Configuration struct {
	DB_NAME     string
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
}

var Db *sql.DB       //Global DB variable for accessing database
var salt = "saltOne" //! Hard coded salt value change this later

// in case of further help is need https://github.com/golangbot/mysqltutorial/blob/master/select/main.go
func EstablishConnection() {

	//Uses the config object to define a mysql config entity
	cfg := mysql.Config{
		User:   config.DB_USERNAME,
		Passwd: config.DB_PASSWORD,
		Net:    "tcp",
		Addr:   config.DB_HOST,
		DBName: config.DB_NAME,
	}

	fmt.Println("Now connecting...")

	//Establishes connection
	Db, _ = sql.Open("mysql", cfg.FormatDSN())

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := Db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("connected...")
}

/*
Function checkExisting will check if a table contains an entity
identical to a passed in one using an id or another identifier
*/
func CheckExisting(tableName string, columnName string, entityID string) bool {
	//Queries the database
	rows, err := Db.Query("SELECT * FROM "+tableName+" WHERE "+columnName+" = ?", entityID)
	if err != nil {
		log.Print("Error when querying" + err.Error())
		return false
	}

	//If the entity already exists
	if rows != nil {
		return true //Return true
	} else { //If the entity doesnt exist
		return false //Return false
	}

}

/*
* Function createNewUser will insert a new email and password into the email and password table
* And create a new credentials entity using the email and password
 */
func CreateNewUser(w http.ResponseWriter, newEmail string, newPassword string) int {
	//Because of DB structure this function is structured very specifically

	//Starting the DB transaction
	ctx := context.Background()
	tx, err := Db.BeginTx(ctx, nil) //Reason being if even one of the three Queries crash none go through
	if err != nil {
		http.Error(w, "Erro starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return 0
	}

	_, execErr := tx.Exec("INSERT INTO `Emails`(`Email`) VALUES (?)", newEmail) //Firstly we need to insert the email since everything is connected to it
	if execErr != nil {
		_ = tx.Rollback()
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(execErr.Error())
		return 0
	}

	password := Hashingsalted(newPassword, salt)

	//Then we create the credentials table because it is connected to the email
	credentials, execErr := tx.Exec("INSERT INTO `Credentials` set `Email`=(SELECT `Email` FROM `Emails` WHERE `Email`=?) , `Password`=? ;", newEmail, password)
	if execErr != nil {
		_ = tx.Rollback()
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(execErr.Error())
		return 0
	}

	//Lastly we create the password which relies on the Credentials table
	_, execErr = tx.Exec("INSERT INTO `Passwords` set `Password`=(SELECT `Password` FROM `Credentials` WHERE `Password`=?) ;", password)
	if execErr != nil {
		_ = tx.Rollback()
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(execErr.Error())
		return 0
	}

	//If all queries go through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when creating new System Manager, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return 0
	}

	id, err := credentials.LastInsertId()
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Println(err.Error())
		return 0
	}
	return int(id)
}
