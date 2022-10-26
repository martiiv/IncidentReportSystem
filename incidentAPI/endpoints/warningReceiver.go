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
)

/*
Class warningReceivers, will handle
- receiver creation
-
Author Aleksander Aaboen
Last rev 25.10 Martin Iversen
*/

/*
Function handleRequest will forward the request to an appropriate function based on method and url
*/
func HandleRequestWarningReceiver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	method := r.Method
	switch method {
	case "GET":
		if r.URL.RawQuery != "" {
			if ok, id := apitools.CorrectQuery(r.URL.RawQuery); ok {
				getWarningReceiver(w, r, id)
			}
		} else {
			getWarningReceivers(w, r)
		}
	case "POST":
		createReceiver(w, r)
	case "DELETE":
		deleteReceiver(w, r)
	}
}

func getWarningReceiver(w http.ResponseWriter, r *http.Request, id string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var warning structs.GetWarningReceiver
	row := databasefunctions.Db.QueryRow("SELECT * FROM `WarningReceiver` WHERE WriD= ?", id)
	if err := row.Scan(
		&warning.Id,
		&warning.Name,
		&warning.PhoneNumber,
		&warning.Company,
		&warning.ReceiverGroup,
		&warning.ReceiverEmail,
	); err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err := json.NewEncoder(w).Encode(warning)
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

}

func getWarningReceivers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	var warnings []structs.GetWarningReceiver
	row, err := databasefunctions.Db.Query("SELECT * FROM `WarningReceiver`")
	if err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for row.Next() {
		warning := structs.GetWarningReceiver{}
		if err = row.Scan(
			&warning.Id,
			&warning.Name,
			&warning.PhoneNumber,
			&warning.Company,
			&warning.ReceiverGroup,
			&warning.ReceiverEmail,
		); err != nil {
			http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
			log.Println(err.Error())
			return
		}

		warnings = append(warnings, warning)
	}
	if err := row.Err(); err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(warnings)
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}
}

func createReceiver(w http.ResponseWriter, r *http.Request) {
	var warningReceiver structs.CreateWarningReceiver
	err := json.NewDecoder(r.Body).Decode(&warningReceiver) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	if !(databasefunctions.CheckExisting("ReceiverGroups", "Name", warningReceiver.ReceiverGroup)) {
		http.Error(w, apitools.UnexpectedError, http.StatusNotImplemented)
		fmt.Fprintf(w, "This receiverGroup does not exist, please use an existing group or create a new one!")
		return
	}

	_, err = databasefunctions.Db.Exec("INSERT INTO `Emails`(`Email`) VALUES (?)", warningReceiver.ReceiverEmail)
	if err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusServiceUnavailable)
		log.Println(err.Error())
		return
	}

	result, err := databasefunctions.Db.Exec("INSERT INTO `WarningReceiver`(`Name`, `PhoneNumber`, `Company`, `ReceiverGroup`, `ReceiverEmail`) VALUES (?,?,?,?,?)", warningReceiver.Name, warningReceiver.PhoneNumber, warningReceiver.Company, warningReceiver.ReceiverGroup, warningReceiver.ReceiverEmail)
	if err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusServiceUnavailable)
		log.Println(err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusServiceUnavailable)
		log.Println(err.Error())
		return
	}

	http.Error(w, "added with "+string(id), http.StatusOK)
}

func deleteReceiver(w http.ResponseWriter, r *http.Request) {
	var warningReceiver structs.DeleteWarningReceiver
	err := json.NewDecoder(r.Body).Decode(&warningReceiver) //Decodes the requests body into the structure defined above
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
	// `tx` is an instance of `*sql.Tx` through which we can execute our queries

	for i := 0; i < len(warningReceiver); i++ {
		// Here, the query is executed on the transaction instance, and not applied to the database yet
		_, err = tx.ExecContext(ctx, "DELETE FROM `WarningReceiver` WHERE WriD = ?", warningReceiver[i].Id)
		if err != nil {
			// Incase we find any error in the query execution, rollback the transaction
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
		http.Error(w, apitools.EncodeError, http.StatusServiceUnavailable)
		log.Println(err.Error())
		return
	}

	wrId := fmt.Sprintf("%#v", warningReceiver)

	http.Error(w, "Successfully deleted Warning receiver with id "+wrId, http.StatusOK)
}
