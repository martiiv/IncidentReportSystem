package databasefunctions

/*
File connection.go, contains basic database functionality for the software

TODO Implement PUT Request
TODO Implement DELETE Request
TODO Error handle
*/
import (
	"database/sql"
	_ "errors"
	"fmt"
	"incidentAPI/config"
	"log"

	"github.com/go-sql-driver/mysql"
)

// Struct Configuration used for getting environment variables for database
type Configuration struct {
	DB_NAME     string
	DB_HOST     string
	DB_USERNAME string
	DB_PASSWORD string
}

var Db *sql.DB

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

	//defer Db.Close()

	// for later testing insrt(db)
	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := Db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("connected...")
}

/*
*
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
*Function createNewUser will insert a new email and password into the email and password table
* And create a new credentials entity using the email and password
 */
func CreateNewUser(newEmail string, newPassword string) int {
	Db.Begin()
	_, err := Db.Exec("INSERT INTO `Emails`(`Email`) VALUES (?)", newEmail)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = Db.Exec("INSERT INTO `Passwords`(`Password`) VALUES (?);", newPassword)
	if err != nil {
		log.Fatal(err.Error())
	}

	credentials, err := Db.Exec("INSERT INTO `Credentials`(`Email`, `Password`) VALUES(?, ?)", newEmail, newPassword)
	if err != nil {
		log.Fatal(err.Error())
	}

	id, err := credentials.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Successfully created credential row with id: %v", int(id))

	return int(id)
}
