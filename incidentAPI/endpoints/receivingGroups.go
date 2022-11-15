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
*File receivingGroups.go will handle all requests related to the receiving group endpoint
? Last revision Martin Iversen 15.11.2022
*/

// Handling the request and forwarding it to the appropriate method
func HandleReceivingGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	url := r.URL.String()
	method := r.Method

	switch method {
	case "GET":
		getReceivingGroups(w, r, url)

	case "POST":
		createReceivingGroups(w, r, url)

	case "PUT":

	case "DELETE":
		deleteReceivingGroups(w, r, url)
	}
}

func getReceivingGroups(w http.ResponseWriter, r *http.Request, url string) {
	variables := mux.Vars(r)
	id := variables["id"]

	if variables["id"] == "" {
		getAllReceivingGroups(w, r) //If the url doesnt contain an id: /incident we want to return all the system managers in the table

	} else if variables["id"] != "" {
		getOneReceivingGroup(w, r, id) //If the url contains an id: /incident?ìd=3 we want to return a spesific systemManager
	} else {
		json.NewEncoder(w).Encode("Please send in an acceptable endpoint URL!")
	}
}

func getAllReceivingGroups(w http.ResponseWriter, r *http.Request) {
	var rcList []structs.GetReceivingGroups

	rows, err := databasefunctions.Db.Query("SELECT `GroupId`, `Name`, `Info` FROM `ReceiverGroups`")
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		receiverGroup := structs.GetReceivingGroups{}
		err = rows.Scan(&receiverGroup.Id, &receiverGroup.Name, &receiverGroup.Info)
		if err != nil {
			http.Error(w, apitools.EncodeError, http.StatusBadRequest)
			log.Fatal(err.Error())
			return
		}

		rcList = append(rcList, receiverGroup)
	}
	rows.Close()

	json.NewEncoder(w).Encode(rcList) //Sends the defined list as a response
}

func getOneReceivingGroup(w http.ResponseWriter, r *http.Request, id string) {
	var receiverGroup structs.GetReceivingGroups

	rows, err := databasefunctions.Db.Query("SELECT `GroupId`, `Name`, `Info` FROM `ReceiverGroups` WHERE `Groupid` = ?", id)
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		err = rows.Scan(&receiverGroup.Id, &receiverGroup.Name, &receiverGroup.Info)
		if err != nil {
			http.Error(w, apitools.EncodeError, http.StatusBadRequest)
			log.Fatal(err.Error())
			return
		}
	}
	rows.Close()

	json.NewEncoder(w).Encode(receiverGroup) //Sends the defined list as a response
}

// Function creates a new receiving group
func createReceivingGroups(w http.ResponseWriter, r *http.Request, url string) {
	var receiverGroup structs.CreateReceivingGroup
	var receiverList []string

	err := json.NewDecoder(r.Body).Decode(&receiverGroup) //Decodes request
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	receiverList = append(receiverList, receiverGroup.Name, receiverGroup.Info)
	databasefunctions.Insrt(w, "ReceiverGroups", receiverList) //Inserts the group into the database

	w.WriteHeader(201)
	fmt.Fprintf(w, "New incident added with name: %v", receiverGroup.Name)
}

// Function deletes a receiving group
func deleteReceivingGroups(w http.ResponseWriter, r *http.Request, url string) {
	var receivingGroup structs.DeleteReceivingGroup
	err := json.NewDecoder(r.Body).Decode(&receivingGroup) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	//For each of receiverGroup struct objects passed in
	for i := 0; i < len(receivingGroup); i++ {
		//If the id field is left as an empty string this means the function should delete based on name instead
		var passingGroup []string
		passingGroup = append(passingGroup, receivingGroup[i].Id)
		passingGroup = append(passingGroup, receivingGroup[i].Name)
		databasefunctions.Delete(w, "ReceiverGroups", passingGroup)
	}
}
