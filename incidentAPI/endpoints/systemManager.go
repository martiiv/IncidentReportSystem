package endpoints

import (
	"context"
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
*/

// Handler for forwarding request to appropriate function based on HTTP method
func HandleSystemManagerRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

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

	rows, err := databasefunctions.Db.Query("SELECT * FROM `SystemManager`")
	if err != nil {
		fmt.Fprintf(w, "Error occurred when querying database, error: %v", err.Error())
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		systemManager := structs.GetSystemManager{}
		err = rows.Scan(
			&systemManager.Id,
			&systemManager.UserName,
			&systemManager.Company,
			&systemManager.Credential)
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

	systemManager := structs.GetSystemManager{}

	rows, err := databasefunctions.Db.Query("SELECT * FROM `SystemManager` WHERE SMiD = ?", id)
	if err != nil {
		fmt.Fprintf(w, "Error occurred when querying database, error: %v", err.Error())
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		systemManager = structs.GetSystemManager{}
		err = rows.Scan(
			&systemManager.Id,
			&systemManager.UserName,
			&systemManager.Company,
			&systemManager.Credential)
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
		log.Println(err.Error())
		return
	}

	//Creates email, credentials and password
	credentialId := databasefunctions.CreateNewUser(w, systemManager.Email, systemManager.Password)

	//Creates the system manager
	object, err := databasefunctions.Db.Exec("INSERT INTO `SystemManager` set `UserName`=? ,`Company`=? ,`Credential`= (SELECT `CiD` FROM `Credentials` WHERE `Cid`=?) ;", systemManager.UserName, systemManager.Company, credentialId)
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	//Fetches the id of the systemManager
	_, err = object.LastInsertId()
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Println(err.Error())
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
		log.Println(err.Error())
		return
	}

	// Create a new context, and begin a transaction
	ctx := context.Background()
	tx, err := databasefunctions.Db.BeginTx(ctx, nil)
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusServiceUnavailable)
		log.Println(err.Error())
		return
	}

	//For each of receiverGroup struct objects passed in
	for i := 0; i < len(systemManager); i++ {

		//Since there is a foreign key constraint on Emails we simply delete the email and the systemManager will be deleted along with the credentials and the password
		_, err = tx.ExecContext(ctx, "DELETE FROM `Emails` WHERE `Email` =?", systemManager[i].Email)
		if err != nil {
			// In case we find any error in the query execution, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
				log.Println(rollbackErr.Error())

				return
			}

			http.Error(w, apitools.EncodeError, http.StatusServiceUnavailable)
			log.Println(err.Error())
			return
		}
	}

	// Finally, if no errors are recieved from the queries, commit the transaction
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusServiceUnavailable)
		log.Println(err.Error())
		return
	}

	wrId := fmt.Sprintf("%#v", systemManager)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Successfully deleted Receiver group with id "+wrId, http.StatusOK)
}
