package endpoints

import (
	"encoding/json"
	"fmt"
	apitools "incidentAPI/apiTools"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Class systemManager will handle all communication with the SystemManager entity in the database
Will handle GET, POST and DELETE
TODO Implement DELETE SystemManager
TODO Error handle
*/

// Handler for forwarding request to appropriate function based on HTTP method
func HandleSystemManagerRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	method := r.Method

	switch method {
	case "GET":
		getSystemManager(w, r, url)

	case "POST":
		createSystemManagers(w, r, url)

	case "PUT":

	case "DELETE":
		deleteSystemManagers(w, r, url)
	}

}

/*
 * Function getSystemManager will forward the GET request based on wether or not the user sends in an id or not
 */
func getSystemManager(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	variables := mux.Vars(r)
	id := variables["id"]

	if variables["id"] == "" {
		getAllSystemManagers(w, r) //If the url doesnt contain an id: /incident we want to return all the system managers in the table

	} else if variables["id"] != "" {
		getOneSystemManager(w, r, id) //If the url contains an id: /incident?ìd=3 we want to return a spesific systemManager
	} else {
		json.NewEncoder(w).Encode("Please send in an acceptable endpoint URL!")
	}
}

/*
* Function fetches al system managers in the database and returns them to the user
 */
func getAllSystemManagers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var smList []structs.GetSystemManager

	rows, err := databasefunctions.Db.Query("SELECT * FROM SystemManager")
	if err != nil {
		fmt.Fprintf(w, "Error occurred when querying database, error: %v", err.Error())
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		systemManager := structs.GetSystemManager{}
		err = rows.Scan(
			&systemManager.Id,
			&systemManager.UserName,
			&systemManager.Company,
			&systemManager.Credentials)
		if err != nil {
			log.Fatal(err)
		}

		smList = append(smList, systemManager)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	json.NewEncoder(w).Encode(smList) //Sends the defined list as a response

}

/*
*Function returns a specific system manager accosiated with the passed in id
 */
func getOneSystemManager(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	systemManager := structs.GetSystemManager{}

	rows, err := databasefunctions.Db.Query("SELECT * FROM `SystemManager`WHERE SMiD = ?", id)
	if err != nil {
		fmt.Fprintf(w, "Error occurred when querying database, error: %v", err.Error())
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		err = rows.Scan(
			&systemManager.Id,
			&systemManager.UserName,
			&systemManager.Company,
			&systemManager.Credentials)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	json.NewEncoder(w).Encode(systemManager) //Sends the defined list as a response
}

func createSystemManagers(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var systemManager structs.CreateSystemManager
	err := json.NewDecoder(r.Body).Decode(&systemManager)
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
	}

	credentialId := databasefunctions.CreateNewUser(systemManager.Email, systemManager.Password)

	object, err := databasefunctions.Db.Exec("INSERT INTO `SystemManager`(`UserName`,`Company`,`Credential`) VALUES (?,?,?)",
		systemManager.UserName,
		systemManager.Company,
		credentialId)

	if err != nil {
		log.Fatal(err)
	}

	id, err := object.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "New System manager added with id: %v", id)
}

func deleteSystemManagers(w http.ResponseWriter, r *http.Request, url string) {

}
