package endpoints

import (
	"encoding/json"
	apitools "incidentAPI/apiTools"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"log"
	"net/http"
)

func LoginSystemManager(w http.ResponseWriter, r *http.Request) {
	var loginStruct structs.LoginStructs

	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	result := databasefunctions.Passwdcheck(w, loginStruct.Password, loginStruct.Email)
	if result != 1 {
		json.NewEncoder(w).Encode("1")
		return
	} else {
		json.NewEncoder(w).Encode("0")
		return
	}

}
