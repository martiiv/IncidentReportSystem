package endpoints

import (
	"encoding/json"
	apitools "incidentAPI/apiTools"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"log"
	"net/http"
)

/*
*File logHandling, will create log reports and similar functionality
! THIS FILE IS PoC ONLY AS WE DONT HAVE TIME
! Missing functionality:
! - Creating a log report with incidents and one big lessons learned from those incidents
! - Creating a report containing all incidents sent to one warning receiver
? Last revision Martin Iversen 22.11.2022
*/

// Handler forwarding requests
func HandleLogRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	method := r.Method

	switch method {
	case "GET":
		returnAccosiatedCountermeasure(w, r)

	case "POST":

	case "PUT":

	case "DELETE":
	}
}

func returnAccosiatedCountermeasure(w http.ResponseWriter, r *http.Request) {
	var accTag structs.TagsStruct
	err := json.NewDecoder(r.Body).Decode(&accTag)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	results, err := databasefunctions.Db.Query("SELECT `Countermeasure` FROM `Incident` WHERE `Tag` = ?", accTag.Tag)
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatalf(err.Error())
		return
	}

	defer results.Close()

	for results.Next() {

	}
}
