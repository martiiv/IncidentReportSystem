package endpoints

import (
	"context"
	"encoding/json"
	apitools "incidentAPI/apiTools"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
* File predefinedCountermeasure
* Used for defining and handling predefined countermeasures for incidents
TODO Implement PUT and DELETE
? Last revision Martin Iversen 21.11.2022
*/

func HandlePDC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	url := r.URL.String()
	method := r.Method

	//Switch case on http methods
	switch method {
	case "GET":
		getPDTCounterM(w, r, url)

	case "POST":
		createPDCounterM(w, r, url)

	case "PUT":
		updatePDCounterM(w, r, url)

	case "DELETE":
		deletePDCounterM(w, r, url)
	}
}

func getPDTCounterM(w http.ResponseWriter, r *http.Request, url string) {
	variables := mux.Vars(r)
	id := variables["tag"]

	if id == "" {
		getAllCounterM(w, r) //If the url doesnt contain an id: /incident we want to return all the predefined countermeasures in the table

	} else if id != "" {
		getOneCounterM(w, r, id) //If the url contains an id: /incident?ìd=3 we want to return a spesific countermeasures
	} else {
		json.NewEncoder(w).Encode("Please send in an acceptable endpoint URL!")
	}
}

func getAllCounterM(w http.ResponseWriter, r *http.Request) {
	var cmList []structs.PDCountermeasure

	results, queryError := databasefunctions.Db.Query("SELECT acTag, Description FROM PredefinedCounterMeasures")
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}
	defer results.Close()

	for results.Next() { //Iterating through results
		var countermeasure structs.PDCountermeasure //Defining ONE incident struct

		if err := results.Scan(&countermeasure.AcTag, &countermeasure.Description); err != nil {
			http.Error(w, "Error scanning results from DB", http.StatusBadRequest)
			log.Fatal(err.Error())
			return
		}

		cmList = append(cmList, countermeasure)
	}

	//All transactions done, we error handle the result
	if err := results.Err(); err != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(cmList) //Sends the defined list as a response

}

func getOneCounterM(w http.ResponseWriter, r *http.Request, tag string) {
	var countermeasure structs.PDCountermeasure

	results, queryError := databasefunctions.Db.Query("SELECT acTag, Description FROM PredefinedCounterMeasures WHERE acTag=?", tag)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}
	defer results.Close()

	for results.Next() { //Iterating through results
		if err := results.Scan(&countermeasure.AcTag, &countermeasure.Description); err != nil {
			http.Error(w, "Error scanning results from DB", http.StatusBadRequest)
			log.Fatal(err.Error())
			return
		}
	}

	//All transactions done, we error handle the result
	if err := results.Err(); err != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Println(err.Error())
		return
	}
	json.NewEncoder(w).Encode(countermeasure) //Sends the defined list as a response
}

func createPDCounterM(w http.ResponseWriter, r *http.Request, url string) {
	var countermeasure structs.PDCountermeasure

	err := json.NewDecoder(r.Body).Decode(&countermeasure)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}
	ctx := context.Background() //Defining context for transaction integration

	tx, err := databasefunctions.Db.BeginTx(ctx, nil) //Start DB transaction
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return
	}

	_, queryError := tx.Exec("INSERT INTO Tags SET acTag = ?")
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}

	_, queryError = tx.Exec("INSERT INTO `PredefinedCounterMeasures` set acTag=(SELECT Tag FROM Tags WHERE Tag = ?), Description= ?", countermeasure.AcTag, countermeasure.Description)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}

	//If query goes through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when inserting rows, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	json.NewEncoder(w).Encode(countermeasure)
}

func updatePDCounterM(w http.ResponseWriter, r *http.Request, url string) {
	var countermeasure structs.PDCountermeasure

	err := json.NewDecoder(r.Body).Decode(&countermeasure)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	_, queryError := databasefunctions.Db.Exec("ALTER TABLE `PredefinedCounterMeasures` set Description= ? WHERE acTag=?", countermeasure.Description, countermeasure.AcTag)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}

	json.NewEncoder(w).Encode("Succsessfully updated countermeasure with tag " + countermeasure.AcTag)
}

func deletePDCounterM(w http.ResponseWriter, r *http.Request, url string) {

}
