package databasefunctions

import (
	"encoding/json"
	_ "errors"
	"fmt"
	apitools "incidentAPI/apiTools"
	"incidentAPI/structs"
	"log"
	"net/http"
)

func Insrt(w http.ResponseWriter, tblname string, params []string) { //if more values are needed adjust accordingly
	var statementtext = "insert into"

	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		stmt, err := Db.Prepare(statementtext + " " + tblname + " set Tag=?, Name= ? , Description= ? , Company= ? , Receiving_group = (SELECT Groupid FROM ReceiverGroups WHERE Name = ?) , Countermeasure = ? , Sendbymanager=(SELECT Username FROM SystemManager WHERE Username = ?) ;")
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return
		}

		_, queryError := stmt.Exec(params[0], params[1], params[2], params[3], params[4], params[5], params[6])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())

			return
		}

	case "WarningReceiver": //this is the case for the table of the Incidents
		stmt, err := Db.Prepare(statementtext + " " + tblname + " set Name= ? , PhoneNumber= ? , Company= ? , ReceiverGroup = (SELECT Name FROM ReceiverGroups WHERE Name = ?) , ReceiverEmail = (SELECT Email FROM Emails WHERE Email = ?) ;")
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return

		}

		_, queryError := stmt.Exec(params[0], params[1], params[2], params[3], params[4])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

		}

	case "ReceiverGroups":
		stmt, err := Db.Prepare(statementtext + " " + tblname + " set Name= ? , Info= ? ;")
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return

		}

		_, queryError := stmt.Exec(params[0], params[1])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return
		}

	case "Emails":
		stmt, err := Db.Prepare(statementtext + " " + tblname + " set Email= ?;")
		if err != nil {
			fmt.Print("helper_methods.go : 118")
			fmt.Println(err)

		}

		_, queryError := stmt.Exec(params[0])
		if queryError != nil {
			fmt.Print("Something went wrong with the execution of the query")
			fmt.Println(queryError)

		}
	}
}

// Delete method it can be adjusted to all the tables and parameters needed this is just if needed in the prototype
func Delete(w http.ResponseWriter, tblname string, params []string) {
	var statementtext = "delete from"

	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		if params[0] == "" {
			stmt, err := Db.Prepare(statementtext + " " + tblname + " where Name=? ;")
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return
			}

			_, queryError := stmt.Exec(params[0], params[1])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())

				return
			}
		} else {
			stmt, err := Db.Prepare(statementtext + " " + tblname + " where IncidentId=? ;")
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return
			}

			_, queryError := stmt.Exec(params[0])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())

				return
			}

		}

	case "ReceiverGroups": //this is the case for the table of the Incidents
		if params[0] == "" {
			stmt, err := Db.Prepare(statementtext + " " + "Emails" + " where Email=? ;")
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return

			}
			_, queryError := stmt.Exec(params[1])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())

				return

			}
		} else {
			stmt, err := Db.Prepare(statementtext + " " + tblname + " where Groupid=? ;")
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return
			}

			_, queryError := stmt.Exec(params[1])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())

				return

			}

		}

	case "WarningReceiver": //this is the case for the table of the Incidents
		if params[0] == "" {
			stmt, err := Db.Prepare(statementtext + " " + tblname + " where Name=? ;")
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return

			}

			_, queryError := stmt.Exec(params[1])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())

				return

			}
		} else {
			stmt, err := Db.Prepare(statementtext + " " + tblname + " where WriD=? ;")
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return

			}

			_, queryError := stmt.Exec(params[0])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())

				return

			}
		}
	}
}

// Function to be used by the manager in case data needs to be altered
func Update(w http.ResponseWriter, tblname string, params []string) {

	var statementtext = "update "

	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		stmt, err := Db.Prepare(statementtext + " " + tblname + " set Name= ? , Context= ? , Company= ? , Credential= ? ,Receiving_group = (SELECT Groupid FROM RecieverGroups WHERE Name = ?) , Countermeasure = ? , Sendbymanager=(SELECT Username FROM SystemManager WHERE Username = ?) " + "Where IncidentId=? ;")
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return
		}

		_, queryError := stmt.Exec(params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())

			return
		}
	}
}

func IncidentSelect(w http.ResponseWriter, incidentId string) {
	var incidentList []structs.GetIncident
	var statementtext = "select "

	if incidentId == "" {
		//According to the name of the table we go to the corresponding action and create the appropriate query
		stmt, err := Db.Prepare(statementtext + " " + "IncidentId, Tag, Name, Description, Company, Receiving_group, Countermeasure, Sendbymanager, Date" + " from Incident ;")
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return
		}

		results, queryError := stmt.Query()
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())

			return
		}

		defer results.Close()

		for results.Next() {
			var incident structs.GetIncident
			if err := results.Scan(&incident.IncidentId, &incident.Tag, &incident.Name, &incident.Description, &incident.Company, &incident.ReceivingGroup, &incident.Countermeasure, &incident.Sendbymanager, &incident.Date); err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(err.Error())

				return
			}

			incidentList = append(incidentList, incident)
		}

		if err := results.Err(); err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return
		}

		json.NewEncoder(w).Encode(incidentList) //Sends the defined list as a response
	} else {

		stmt, err := Db.Prepare(statementtext + " " + "IncidentId, Tag, Name, Description, Company, Receiving_group, Countermeasure, Sendbymanager, Date" + " from Incident WHERE `IncidentId` = " + incidentId + " ;")
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())

			return
		}

		results, queryError := stmt.Query()
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())

			return
		}

		defer results.Close()

		incident := structs.GetIncident{}

		for results.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct

			err = results.Scan(&incident.IncidentId, &incident.Tag, &incident.Name, &incident.Description, &incident.Company, &incident.ReceivingGroup, &incident.Countermeasure, &incident.Sendbymanager, &incident.Date)
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Println(queryError.Error())
			}
		}

		err = results.Err()
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())
		}

		results.Close()
		json.NewEncoder(w).Encode(incident) //Sends the defined list as a response
	}
}

func Select_warning_receivers(w http.ResponseWriter) {
	var statementtext = "select "
	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	stmt, err := Db.Prepare(statementtext + " " + "WriD , Name, PhoneNumber , Company , ReceiverGroup , ReceiverId " + " from WarningReceiver ;")
	if err != nil {
		fmt.Print("helper_methods.go : 118")
		fmt.Println(err)

	}
	results, queryError := stmt.Query()
	if queryError != nil {
		fmt.Print("Something went wrong with the execution of the query")
		fmt.Println(queryError)

	}

	defer results.Close()
	fmt.Println("Results from select query: ")

	for results.Next() {
		var datares [6]string

		if err := results.Scan(&datares[0], &datares[1], &datares[2], &datares[3], &datares[4], &datares[5]); err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s\n", datares)

	}

	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}

func SelecTags(w http.ResponseWriter) {
	var tags []string
	var statementtext = "SELECT "
	stmt, err := Db.Prepare(statementtext + " " + " `Tag` FROM `Incident` GROUP BY `Tag")
	if err != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Println(err.Error())

		return
	}

	results, queryError := stmt.Query()
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Println(err.Error())

		return
	}

	for results.Next() {
		var dbResponse string
		if err := results.Scan(&dbResponse); err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())
		}
		tags = append(tags, dbResponse)
	}
	json.NewEncoder(w).Encode(tags)
}
