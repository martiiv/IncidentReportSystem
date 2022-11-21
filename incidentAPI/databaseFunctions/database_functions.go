package databasefunctions

import (
	"context"
	"encoding/json"
	_ "errors"
	"fmt"
	apitools "incidentAPI/apiTools"
	"incidentAPI/structs"
	"log"
	"net/http"
)

/*
* File database_functions
* Handles all generic transactions
? Last revision Martin iversen 15.11.2022
*/

// Function Insrt used when inserting data into the different tables in the database
func Insrt(w http.ResponseWriter, tblname string, params []string) {
	var statementtext = "INSERT INTO " //Predefined statement for inserting data
	ctx := context.Background()        //Defining context for transaction integration
	tx, err := Db.BeginTx(ctx, nil)    //Start DB transaction
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return
	}

	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {

	case "Incident": //this is the case for the table of the Incidents
		//Executing query with parameters
		_, queryError := tx.Exec(statementtext+"`Incident` set Tag=?, Name= ? , Description= ? , Company= ? , Receiving_group = (SELECT Groupid FROM ReceiverGroups WHERE Name = ?) , Countermeasure = ? , Sendbymanager=(SELECT Username FROM SystemManager WHERE Username = ?) ;", params[0], params[1], params[2], params[3], params[4], params[5], params[6])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}

	case "WarningReceiver": //this is the case for the table of the Incidents
		//Executing query
		_, queryError := tx.Exec(statementtext+" "+tblname+" set Name= ? , PhoneNumber= ? , Company= ? , ReceiverGroup = (SELECT Name FROM ReceiverGroups WHERE Name = ?) , ReceiverEmail = (SELECT Email FROM Emails WHERE Email = ?) ;", params[0], params[1], params[2], params[3], params[4])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}

	case "ReceiverGroups":
		//Executing query
		_, queryError := tx.Exec(statementtext+" "+tblname+" set Name= ? , Info= ? ;", params[0], params[1])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())
			return
		}

	case "Tags":
		_, queryError := tx.Exec(statementtext+" "+tblname+" set Tag= ?;", params[0])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())
			return
		}
	case "PredefinedCounterMeasures":
		_, queryError := tx.Exec(statementtext+" "+tblname+" set acTag= (SELECT Tag FROM Tags WHERE Tag = ?), Description = ?;", params[0], params[1])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(queryError.Error())
			return
		}
	}
	//If query goes through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when inserting rows, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}
// Delete method it can be adjusted to all the tables and parameters needed this is just if needed in the prototype
func Delete(w http.ResponseWriter, tblname string, params []string) {
	var statementtext = "delete from"
	ctx := context.Background()     //Defining context for transaction integration
	tx, err := Db.BeginTx(ctx, nil) //Start DB transaction
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return
	}

	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {

	case "Incident": //this is the case for the table of the Incidents
		if params[0] == "" { //If the id field is empty
			_, queryError := tx.Exec(statementtext+" "+tblname+" where Name=? ;", params[1]) //We delete based on name of incident
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(queryError.Error())
				return
			}

		} else { //If the id field is defined
			_, queryError := tx.Exec(statementtext+" "+tblname+" where IncidentId=? ;", params[0]) //We delete based on incidentId
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(queryError.Error())

				return
			}
		}

	case "ReceiverGroups": //this is the case for the table of the Incidents
		if params[0] == "" { //If Id is empty
			_, queryError := tx.Exec(statementtext+" `ReceiverGroups`"+" where Name = ?;", params[1]) //We delete based on name
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Fatal(queryError.Error())
				return
			}

		} else {
			_, queryError := tx.Exec(statementtext+" "+tblname+" where Groupid=? ;", params[0])
			if queryError != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Fatal(queryError.Error())
				return
			}
		}

	case "WarningReceiver": //this is the case for the table of the Incidents
		_, queryError := tx.Exec(statementtext+" "+"`Emails`"+" where `Email`= ? ;", params[1])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}
	}

	//If query goes through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when deleting rows, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

// Function to be used by the manager in case data needs to be altered
func Update(w http.ResponseWriter, tblname string, params []string) {
	var statementtext = "update "
	ctx := context.Background()     //Defining context for transaction integration
	tx, err := Db.BeginTx(ctx, nil) //Start DB transaction
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return
	}

	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		_, queryError := tx.Exec(statementtext+" "+tblname+" set Name= ? , Context= ? , Company= ? , Credential= ? ,Receiving_group = (SELECT Groupid FROM RecieverGroups WHERE Name = ?) , Countermeasure = ? , Sendbymanager=(SELECT Username FROM SystemManager WHERE Username = ?) "+"Where IncidentId=? ;", params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}
		//TODO Implement more cases if necessary
	}

	//If query goes through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when deleting rows, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}

//Updates the text in lesson learned row of the table incidents in the DB according to the selected incident based on the ID
func Updatelessonslearned(w http.ResponseWriter,params []string){
	var statementtext = "update "
	ctx := context.Background()     //Defining context for transaction integration
	tx, err := Db.BeginTx(ctx, nil) //Start DB transaction
	if err != nil {
		http.Error(w, "Error starting database transaction", http.StatusBadGateway)
		log.Fatal(err.Error())
		return
	}
	//Query execution
		_, queryError := tx.Exec(statementtext+" "+"Incident"+" set LessonLearned= '?' "+"Where IncidentId=? ;", params[0], params[1])
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}
		//TODO Implement more cases if necessary
	}

	//If query goes through we commit the transactions
	if err := tx.Commit(); err != nil {
		http.Error(w, "Error encountered when deleting rows, rolling back transactions...", http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}


// Function for selecting incidents
// ! Needs rewriting but we dont have time
func IncidentSelect(w http.ResponseWriter, incidentId string) {
	var incidentList []structs.GetIncident //Defines incidentstruct
	var statementtext = "select "          //Defines statement

	if incidentId == "" {
		//According to the name of the table we go to the corresponding action and create the appropriate query
		results, queryError := Db.Query(statementtext + " " + "IncidentId, Tag, Name, Description, Company, Receiving_group, Countermeasure, Sendbymanager, Date" + " from Incident ;")
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}
		defer results.Close()

		for results.Next() { //Iterating through results
			var incident structs.GetIncident //Defining ONE incident struct

			if err := results.Scan(&incident.IncidentId, &incident.Tag, &incident.Name, &incident.Description, &incident.Company, &incident.ReceivingGroup, &incident.Countermeasure, &incident.Sendbymanager, &incident.Date); err != nil {
				http.Error(w, "Error scanning results from DB", http.StatusBadRequest)
				log.Fatal(err.Error())
				return
			}

			//Finding appropriate Group name from ID in incident
			groupsName, err := Db.Query("SELECT `Name` FROM `ReceiverGroups` WHERE `GroupId` = ? ;", incident.ReceivingGroup)
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(err.Error())
				return
			}

			for groupsName.Next() { //Iterating through these results
				var name string
				if err := groupsName.Scan(&name); err != nil {
					http.Error(w, apitools.QueryError, http.StatusInternalServerError)
					log.Fatal(err.Error())
					return
				}
				//Defining the name of the group
				incident.ReceivingGroup = name
				incidentList = append(incidentList, incident) //Adding incident to the list
			}
		}
		//All transactions done, we error handle the result
		if err := results.Err(); err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
		json.NewEncoder(w).Encode(incidentList) //Sends the defined list as a response

	} else { //An ID is passed in we get one specific incident
		incident := structs.GetIncident{}

		results, queryError := Db.Query(statementtext + " " + "IncidentId, Tag, Name, Description, Company, Receiving_group, Countermeasure, Sendbymanager, Date" + " from Incident WHERE `IncidentId` = " + incidentId + " ;")
		if queryError != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}
		defer results.Close()

		for results.Next() { //For all the rows in the database we convert the sql entity to a string and insert it into a struct
			err := results.Scan(&incident.IncidentId, &incident.Tag, &incident.Name, &incident.Description, &incident.Company, &incident.ReceivingGroup, &incident.Countermeasure, &incident.Sendbymanager, &incident.Date)
			if err != nil {
				http.Error(w, apitools.QueryError, http.StatusBadRequest)
				log.Fatal(err.Error())
				return
			}
		}

		groupsName, err := Db.Query("SELECT `Name` FROM `ReceiverGroups` WHERE `GroupId` = ? ;", incident.ReceivingGroup)
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusInternalServerError)
			log.Fatal(err.Error())
			return
		}
		var name string
		for groupsName.Next() {
			if err := groupsName.Scan(&name); err != nil {
				http.Error(w, apitools.QueryError, http.StatusInternalServerError)
				log.Fatal(err.Error())
				return
			}
			incident.ReceivingGroup = name
		}

		err = results.Err()
		if err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(queryError.Error())
			return
		}
		results.Close()

		json.NewEncoder(w).Encode(incident) //Sends the defined list as a response
	}
}

// Function for selecting warning receivers
func Select_warning_receivers(w http.ResponseWriter) {
	var statementtext = "select "
	//Executing query
	results, queryError := Db.Query(statementtext + " " + "WriD , Name, PhoneNumber , Company , ReceiverGroup , ReceiverId " + " from WarningReceiver ;")
	if queryError != nil {
		fmt.Print("Something went wrong with the execution of the query")
		fmt.Println(queryError.Error())
		return

	}
	defer results.Close() //Closing results

	for results.Next() { //Scanning results
		var datares [6]string
		if err := results.Scan(&datares[0], &datares[1], &datares[2], &datares[3], &datares[4], &datares[5]); err != nil {
			http.Error(w, apitools.QueryError, http.StatusInternalServerError)
			log.Fatal(err.Error())
			return
		}
		fmt.Fprintf(w, "%s\n", datares) //Returning response
	}

	if err := results.Err(); err != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}
}

// Function for selecting tags from incidents
func SelecTags(w http.ResponseWriter) {
	var tags []structs.TagsStruct
	var statementtext = "SELECT "

	results, queryError := Db.Query(statementtext + " " + " `Tag` FROM `Incident` GROUP BY `Tag`")
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}

	for results.Next() {
		var dbResponse string
		if err := results.Scan(&dbResponse); err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(err.Error())
		}

		result := structs.TagsStruct{Tag: dbResponse}
		tags = append(tags, result)
	}
	json.NewEncoder(w).Encode(tags)
}

// Function for selecting countermeasures with their tags from Predifined Countermeasures table
func SelecCmeasures(w http.ResponseWriter , tag string) {
	var description []structs.Countermeasure
	var statementtext = "SELECT "
//based on the tag selected from the UI the query will get the appropriate countermeasure
	results, queryError := Db.Query(statementtext + " " + "Description FROM `PredefinedCounterMeasures` where acTag = '?' " , tag)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusBadRequest)
		log.Fatal(queryError.Error())
		return
	}

	for results.Next() {
		var dbResponse string
		if err := results.Scan(&dbResponse); err != nil {
			http.Error(w, apitools.QueryError, http.StatusBadRequest)
			log.Fatal(err.Error())
		}

		result := structs.Countermeasure{Description: dbResponse}
		description = append(description, result)
	}
	json.NewEncoder(w).Encode(description)
}
