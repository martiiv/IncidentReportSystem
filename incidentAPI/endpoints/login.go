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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Request-Method", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var loginStruct structs.LoginStructs

	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	result := databasefunctions.Passwdcheck(w, loginStruct.Password, loginStruct.Email)
	if result != 1 {
		http.Error(w, "Logged in failed", http.StatusUnauthorized)
		return
	} else {
		credentials := loggedIn(loginStruct.Email)
		credentials.Email = loginStruct.Email
		json.NewEncoder(w).Encode(credentials) //Sends the defined list as a response
		return
	}

}

func loggedIn(email string) structs.LoggedIn {

	systemManager := structs.LoggedIn{}

	rows, err := databasefunctions.Db.Query("SELECT SystemManager.Username, SystemManager.Company, SystemManager.Credential FROM `SystemManager` WHERE Credential = (SELECT CiD from Credentials WHERE Email = ?)", email)
	if err != nil {
		return structs.LoggedIn{}
	}

	for rows.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
		err = rows.Scan(
			&systemManager.UserName,
			&systemManager.Company,
			&systemManager.CiD,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	rows.Close()

	return systemManager

}
