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
* File systemManager.go will handle all communication with the SystemManager entity in the database
* Will handle GET, POST and DELETE
? Last revision Martin Iversen 15.11.2022
*/

// Handler for forwarding request to appropriate function based on HTTP method
func HandleSystemManagerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Request-Method", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	url := r.URL.String()
	method := r.Method

	switch method {
	case "GET":
		getSystemManager(w, r, url)

	case "POST":
		createSystemManagers(w, r, url)

	case "PUT":
		LoginSystemManager(w, r)

	case "DELETE":
		deleteSystemManagers(w, r, url)
	}

}

/*
 * Function getSystemManager will forward the GET request based on wether or not the user sends in an id or not
 */
func getSystemManager(w http.ResponseWriter, r *http.Request, url string) {

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
	var smList []structs.GetSystemManager

	rows, err := databasefunctions.Db.Query("SELECT `SMiD`, `Username`, `Company`,`Credential` FROM `SystemManager`")
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		systemManager := structs.GetSystemManager{}
		err = rows.Scan(&systemManager.Id, &systemManager.UserName, &systemManager.Company, &systemManager.Credential)
		if err != nil {
			http.Error(w, apitools.EncodeError, http.StatusBadRequest)
			log.Fatal(err.Error())
			return
		}

		smList = append(smList, systemManager)
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	rows.Close()

	json.NewEncoder(w).Encode(smList) //Sends the defined list as a response
}

/*
*Function returns a specific system manager accosiated with the passed in id
 */
func getOneSystemManager(w http.ResponseWriter, r *http.Request, id string) {

	systemManager := structs.GetSystemManager{}

	rows, err := databasefunctions.Db.Query("SELECT `SMiD`, `Username`, `Company`,`Credential` FROM `SystemManager` WHERE SMiD = ?", id)
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		systemManager = structs.GetSystemManager{}
		err = rows.Scan(&systemManager.Id, &systemManager.UserName, &systemManager.Company, &systemManager.Credential)
		if err != nil {
			http.Error(w, apitools.EncodeError, http.StatusBadRequest)
			log.Fatal(err.Error())
			return
		}
	}

	err = rows.Err()
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	rows.Close()

	json.NewEncoder(w).Encode(systemManager) //Sends the defined list as a response
}

/*
Function for creating system Manager
* Creates an email row
*					->A credentials row
*									->Password row
*												->SystemManager row
* Reason for this workflow is the foreign key constraints it makes deletion alot more efficent
*/
func createSystemManagers(w http.ResponseWriter, r *http.Request, url string) {
	var systemManager structs.CreateSystemManager

	err := json.NewDecoder(r.Body).Decode(&systemManager)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	//Creates email, credentials and password
	credentialId := databasefunctions.CreateNewUser(w, systemManager.Email, systemManager.Password)

	//Creates the system manager
	object, err := databasefunctions.Db.Exec("INSERT INTO `SystemManager` set `UserName`=? ,`Company`=? ,`Credential`= (SELECT `CiD` FROM `Credentials` WHERE `Cid`=?) ;", systemManager.UserName, systemManager.Company, credentialId)
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	//Fetches the id of the systemManager
	_, err = object.LastInsertId()
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated) //Returns the appropriate status code
	fmt.Fprintf(w, "New System manager added with name: %v", systemManager.UserName)
}

/*
Function will delete the system manager
* Deletes the Email
*				-> Deletes the credentials
*										-> Deletes the password
*										-> Deletes the system manager
*/
func deleteSystemManagers(w http.ResponseWriter, r *http.Request, url string) {
	var systemManager structs.DeleteSystemManager
	err := json.NewDecoder(r.Body).Decode(&systemManager) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	//For each of receiverGroup struct objects passed in
	for i := 0; i < len(systemManager); i++ {

		//Since there is a foreign key constraint on Emails we simply delete the email and the systemManager will be deleted along with the credentials and the password
		_, err = databasefunctions.Db.Exec("DELETE FROM `Emails` WHERE `Email` =?", systemManager[i].Email)
		if err != nil {
			http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
			log.Fatal(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Successfully deleted System manager with email "+systemManager[i].Email, http.StatusOK)
	}
}
