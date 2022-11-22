package endpoints

import (
	"encoding/json"
	"fmt"
	apitools "incidentAPI/apiTools"
	"incidentAPI/communication"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
*File incident.go, will handle
- incident creation
- insertion of countermeasures
- fetching incidents from the database

!You can only send in one recieverGroup when creating the incident
? Last revision Martin Iversen 15.11.2022
*/

/*
Function handleRequest will forward the request to an appropriate function based on method and url
*/
func HandleIncidentRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	url := r.URL.String()
	method := r.Method

	//Switch case on http methods
	switch method {
	case "GET":
		getIncident(w, r, url)

	case "POST":
		createIncident(w, r, url)

	case "PUT":
		updateLessonsLearned(w, r, url)

	case "DELETE":
		deleteIncident(w, r, url)
	}

}

/*
Function getIncidents, will fetch all incidents, one specific incident or incidents based on certain parameters
Function does forward the request to the appropriate endpoint based on wether or not the url contains parameters or not
*/
func getIncident(w http.ResponseWriter, r *http.Request, url string) {
	variables := mux.Vars(r)
	id := variables["id"]
	tag := variables["tag"]

	if id == "" && tag == "" {
		getAllIncidents(w, r) //If the url doesnt contain an id: /incident we want to return all the incidents in the table

	} else if id != "" && tag == "" {
		getOneIncident(w, r, id) //If the url contains an id: /incident?ìd=3 we want to return a spesific incident
	} else if tag == "true" {
		getAvailableTags(w, r)
	} else {
		json.NewEncoder(w).Encode("Please send in an acceptable endpoint URL!")
	}
}

/*
Function will fetch all incidents in the Incident table in the database
*/
func getAllIncidents(w http.ResponseWriter, r *http.Request) {
	databasefunctions.IncidentSelect(w, "")
}

/*
Function will fetch one specific incident from the database based on the passed in ID
*/
func getOneIncident(w http.ResponseWriter, r *http.Request, id string) {
	databasefunctions.IncidentSelect(w, id)
}

/*
Function creaetIncident will create a new incident in the database
*/
func createIncident(w http.ResponseWriter, r *http.Request, url string) {
	var incident structs.CreateIncident
	var incidentList []string

	err := json.NewDecoder(r.Body).Decode(&incident)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	//Checks whether or not the Recieving group exists
	checkGroup := databasefunctions.CheckExisting("ReceiverGroups", "Groupid", incident.ReceivingGroup)

	if !checkGroup {
		json.NewEncoder(w).Encode("Group does not exist! Please use an existing group or create one with this name!")
	}

	incidentList = append(incidentList, incident.Tag, incident.Name, incident.Description, incident.Company, incident.ReceivingGroup, incident.Sendbymanager, incident.LessonLearned)
	databasefunctions.Insrt(w, "Incident", incidentList)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New incident added with name: %v", incident.Name)

	//?Check if error
	err = communication.SendMail(w, incident)
	if err != nil {
		http.Error(w, "Error occurred when sending email!", http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}
}

/*
Function updateCountermeasures will update the incidents suggested countermeasures in the database
*/
func updateLessonsLearned(w http.ResponseWriter, r *http.Request, url string) {
	var LessonLearned structs.UpdateLessonsLearned

	err := json.NewDecoder(r.Body).Decode(&LessonLearned)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	_, err = databasefunctions.Db.Exec("UPDATE `Incident` SET `LessonLearned` = ? WHERE `IncidentId` = ?", LessonLearned.LessonLearned, LessonLearned.IncidentId)
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully updated Countermeasures in incident: %v", LessonLearned.IncidentId)
}

/*
Function deleteIncident will delete a given incident from the database
Takes in either the incident id or the incidentName in order to delete an instance of incident from the database
*/
func deleteIncident(w http.ResponseWriter, r *http.Request, url string) {
	var incident structs.DeleteIncident
	err := json.NewDecoder(r.Body).Decode(&incident) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	//For each of incident struct objects passed in
	for i := 0; i < len(incident); i++ {
		var params []string
		params = append(params, incident[i].IncidentId, incident[i].IncidentName)
		databasefunctions.Delete(w, "Incident", params)
	}
}

// Function for return all tags in the database
func getAvailableTags(w http.ResponseWriter, r *http.Request) {
	databasefunctions.SelecTags(w)
}
