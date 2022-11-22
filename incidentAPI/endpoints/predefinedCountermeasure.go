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
* File predefinedCountermeasure
* Used for defining and handling predefined countermeasures for incidents
? Last revision Martin Iversen 22.11.2022
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

// FUnction for getting all countermeasures
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

		if err := results.Scan(&countermeasure.Tag, &countermeasure.Description); err != nil {
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cmList) //Sends the defined list as a response
}

// Function for getting one countermeasure
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
		if err := results.Scan(&countermeasure.Tag, &countermeasure.Description); err != nil {
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(countermeasure) //Sends the defined list as a response
}

// Function for creating new countermeasures
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

	//If the tag already exists we update it instead of creating a duplicate tag/an error gets thrown
	if !databasefunctions.CheckExisting("Tags", "Tag", countermeasure.Tag) {
		_, queryError := tx.Exec("INSERT INTO Tags SET Tag = ?", countermeasure.Tag)
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			tx.Rollback()
			return
		}
	}

	_, queryError := tx.Exec("INSERT INTO `PredefinedCounterMeasures` set acTag=(SELECT Tag FROM Tags WHERE Tag = ?), Description= ? ;", countermeasure.Tag, countermeasure.Description)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		tx.Rollback()
		return
	}

	//If query goes through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when inserting rows, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New countermeasure added with tag: %v", countermeasure.Tag)
}

// Function for updating countermeasures
func updatePDCounterM(w http.ResponseWriter, r *http.Request, url string) {
	var countermeasure structs.PDCountermeasure

	err := json.NewDecoder(r.Body).Decode(&countermeasure)
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	_, queryError := databasefunctions.Db.Exec("UPDATE `PredefinedCounterMeasures` set Description= ? WHERE acTag=?", countermeasure.Description, countermeasure.Tag)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Succsessfully updated countermeasure with tag %v", countermeasure.Tag)
}

// Function for deleteing predefined countermeasures
// ! This wont be implemented due to product logic, for deletion we set the description to an empty string instead
func deletePDCounterM(w http.ResponseWriter, r *http.Request, url string) {
	var countermeasure []structs.DeleteCountermeasure

	err := json.NewDecoder(r.Body).Decode(&countermeasure) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	ctx := context.Background()                       //Defining context for transaction integration
	tx, err := databasefunctions.Db.BeginTx(ctx, nil) //Start DB transaction
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return
	}

	//For each of incident struct objects passed in
	for i := 0; i < len(countermeasure); i++ {
		if countermeasure[i].Cascade == "Yes" {
			_, queryError := tx.Exec("DELETE FROM Tags WHERE Tag = ?", countermeasure[i].Tag)
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(queryError.Error())
				return
			}

			_, queryError = tx.Exec("DELETE FROM PredefinedCounterMeasures WHERE acTag = ?", countermeasure[i].Tag)
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(queryError.Error())
				return
			}

		} else if countermeasure[i].Cascade == "No" {
			_, queryError := tx.Exec("DELETE FROM PredefinedCounterMeasures WHERE acTag = ?", countermeasure[i].Tag)
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(queryError.Error())
				return
			}
		} else {
			fmt.Fprint(w, "Please pass in either Yes or No in the cascade field in the request body!")
			return
		}

		//If query goes through we commit the transactions
		if err := tx.Commit(); err != nil {
			http.Error(w, "Error encountered when deleting rows, rolling back transactions...", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
