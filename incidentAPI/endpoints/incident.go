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

Author Martin Iversen
Last rev 17.10
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
*/
func getIncident(w http.ResponseWriter, r *http.Request, url string) {
	variables := mux.Vars(r)
	id := variables["id"]
	var incident structs.GetAllIncidents

	log.Print(id)

	log.Print("Bro hva faen skjer her")

	//Insert database SQL statement to get incident with ID
	rows, err := databasefunctions.Db.Query("SELECT * FROM Incident WHERE IncidentId=?", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Rett før scan")

	for rows.Next() {
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
	}

	log.Print("Etter scan")

	if err != nil {
		log.Fatal(err)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(incident)
}

/*
Function creaetIncident will create a new incident in the database
*/
func createIncident(w http.ResponseWriter, r *http.Request, url string) {

}

/*
Function updateCountermeasures will update the incidents suggested countermeasures in the database
*/
func updateCountermeasures(w http.ResponseWriter, r *http.Request, url string) {

}

/*
Function deleteIncident will delete a given incident from the database
*/
func deleteIncident(w http.ResponseWriter, r *http.Request, url string) {

}
