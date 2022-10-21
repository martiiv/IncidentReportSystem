package endpoints

import (
	"encoding/json"
	"fmt"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
Class incident, will handle
- incident creation
- insertion of countermeasures
- fetching incidents from the database

TODO create DELETE functionality
!You can only send in one recieverGroup when creating the incident

Author Martin Iversen
Last rev 19.10
*/

/*
Function handleRequest will forward the request to an appropriate function based on method and url
*/
func HandleIncidentRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	method := r.Method

	switch method {
	case "GET":
		getIncident(w, r, url)

	case "POST":
		createIncident(w, r, url)

	case "PUT":
		updateCountermeasures(w, r, url)

	case "DELETE":
		deleteIncident(w, r, url)
	}

}

/*
Function getIncidents, will fetch all incidents, one specific incident or incidents based on certain parameters
Function does forward the request to the appropriate endpoint based on wether or not the url contains parameters or not
*/
func getIncident(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	variables := mux.Vars(r)
	id := variables["id"]

	if variables["id"] == "" {
		getAllIncidents(w, r) //If the url doesnt contain an id: /incident we want to return all the incidents in the table

	} else if variables["id"] != "" {
		getOneIncident(w, r, id) //If the url contains an id: /incident?ìd=3 we want to return a spesific incident
	} else {
		json.NewEncoder(w).Encode("Please send in an acceptable endpoint URL!")
	}
}

/*
Function will fetch all incidents in the Incident table in the database
*/
func getAllIncidents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var incidentList []structs.GetIncident

	rows, err := databasefunctions.Db.Query("SELECT * FROM `Incident`") //Defines the sql request
	if err != nil {
		fmt.Fprintf(w, "Error occurred when querying database, error: %v", err.Error())
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		incident := structs.GetIncident{}
		err = rows.Scan(
			&incident.IncidentId,
			&incident.Tag,
			&incident.Name,
			&incident.Description,
			&incident.Company,
			&incident.Credential,
			&incident.ReceivingGroup,
			&incident.Countermeasure,
			&incident.Sendbymanager,
			&incident.Date)
		if err != nil {
			log.Fatal(err)
		}

		incidentList = append(incidentList, incident)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	json.NewEncoder(w).Encode(incidentList) //Sends the defined list as a response
}

/*
Function will fetch one specific incident from the database based on the passed in ID
*/
func getOneIncident(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	incident := structs.GetIncident{}
	rows, err := databasefunctions.Db.Query("SELECT * FROM `Incident` WHERE `IncidentId` = ?", id) //Defines the SQL statement with ID
	if err != nil {
		fmt.Fprintf(w, "Error occurred when querying database, error: %v", err.Error())
	}

	for rows.Next() {
		err = rows.Scan( //Converts the columns and inserts them into appropriate struct
			&incident.IncidentId,
			&incident.Tag,
			&incident.Name,
			&incident.Description,
			&incident.Company,
			&incident.Credential,
			&incident.ReceivingGroup,
			&incident.Countermeasure,
			&incident.Sendbymanager,
			&incident.Date)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(incident) //Sends the incident object as a response
}

/*
Function creaetIncident will create a new incident in the database
*/
func createIncident(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var incident structs.CreateIncident
	err := json.NewDecoder(r.Body).Decode(&incident)
	if err != nil {
		log.Fatal(err)
		return
	}

	//Checks wether or not the Recieving group exists
	checkGroup := databasefunctions.CheckExisting("ReceiverGroups", "Groupid", incident.ReceivingGroup)

	if !checkGroup {
		//! Currently you can only pass in one recieving group, we need to be able to implement several groups
		json.NewEncoder(w).Encode("Group does not exist! Please use an existing group or create one with this name!")
		log.Print(checkGroup) //Remove this later
	}

	object, err := databasefunctions.Db.Exec("INSERT INTO `Incident` (`Tag`, `Name`, `Description`, `Company`, `Credential`, `Receiving_Group`, `Countermeasure`, `Sendbymanager`, `Date`) VALUES(?,?,?,?,?,?,?,?,?)",
		incident.Tag,
		incident.Name,
		incident.Description,
		incident.Company,
		incident.Credential,
		incident.ReceivingGroup,
		incident.Countermeasure,
		incident.Sendbymanager,
		incident.Date)

	if err != nil {
		log.Fatal(err)
	}

	id, err := object.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "New incident added with id: %v", id)
}

/*
Function updateCountermeasures will update the incidents suggested countermeasures in the database
*/
func updateCountermeasures(w http.ResponseWriter, r *http.Request, url string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var countermeasure structs.UpdateCountermeasure

	err := json.NewDecoder(r.Body).Decode(&countermeasure)
	if err != nil {
		fmt.Fprintf(w, "Error occurred: %v", err.Error())
	}

	rows, err := databasefunctions.Db.Exec("UPDATE `Incident` SET `Countermeasure` = ? WHERE `IncidentId` = ?", countermeasure.Countermeasure, countermeasure.IncidentId)
	if err != nil {
		fmt.Fprintf(w, "Error occurred: %v", err.Error())
	}
	log.Print(rows.RowsAffected()) //TODO Remove this or smthing idk

	fmt.Fprintf(w, "Successfully updated Countermeasures in incident: %v", countermeasure.IncidentId)
}

/*
Function deleteIncident will delete a given incident from the database
*/
func deleteIncident(w http.ResponseWriter, r *http.Request, url string) {

}
