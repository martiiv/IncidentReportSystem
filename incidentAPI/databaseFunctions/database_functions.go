package databasefunctions

import (
	"database/sql"
	_ "errors"
	"fmt"
	"log"
)

func Insrt(dbname *sql.DB, tblname string, params []string) { //if more values are needed adjust accordingly
	var statementtext = "insert into"

	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		stmt, err := dbname.Prepare(statementtext + " " + tblname + " set Tag=?, Name= ? , Context= ? , Company= ? , Credential= ? ,Receiving_group = (SELECT Groupid FROM ReceiverGroups WHERE Name = ?) , Countermeasure = ? , Sendbymanager=(SELECT Username FROM SystemManager WHERE Username = ?) ;")
		if err != nil {
			fmt.Print("helper_methods.go : 118")
			fmt.Println(err)

		}

		_, queryError := stmt.Exec(params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7])
		if queryError != nil {
			fmt.Print("Something went wrong with the execution of the query")
			fmt.Println(queryError)

		}
		fmt.Print("data added in table" + tblname)

	case "WarningReceiver": //this is the case for the table of the Incidents
		stmt, err := dbname.Prepare(statementtext + " " + tblname + " set Name= ? , PhoneNumber= ? , CredentialId= (SELECT CiD FROM Credentials WHERE Email = ?) , Company= ? , ReceiverGroup = (SELECT Name FROM ReceiverGroups WHERE Name = ?) , ReceiverId = (SELECT Groupid FROM ReceiverGroups WHERE Name = ?) ;")
		if err != nil {
			fmt.Print("helper_methods.go : 118")
			fmt.Println(err)

		}

		_, queryError := stmt.Exec(params[0], params[1], params[2], params[3], params[4], params[5])
		if queryError != nil {
			fmt.Print("Something went wrong with the execution of the query")
			fmt.Println(queryError)

		}
		fmt.Print("data added in table" + tblname)
	case "ReceiverGroups":
		stmt, err := dbname.Prepare(statementtext + " " + tblname + " set Name= ? , Info= ? ;")
		if err != nil {
			fmt.Print("helper_methods.go : 118")
			fmt.Println(err)

		}

		_, queryError := stmt.Exec(params[0], params[1])
		if queryError != nil {
			fmt.Print("Something went wrong with the execution of the query")
			fmt.Println(queryError)

		}
		fmt.Print("data added in table" + tblname)


	}


}

// Delete method it can be adjusted to all the tables and parameters needed this is just if needed in the prototype
func Delet(dbname *sql.DB, tblname string, params []string) {
	var statementtext = "delete from"

	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		stmt, err := dbname.Prepare(statementtext + " " + tblname + " where IncidentId=? and Name=? ;")
		if err != nil {
			fmt.Print("helper_methods.go : 118")
			fmt.Println(err)

		}

		_, queryError := stmt.Exec(params[0], params[1])
		if queryError != nil {
			fmt.Print("Something went wrong with the execution of the query")
			fmt.Println(queryError)

		}

	}
}

// Function to be used by the manager in case data needs to be altered
func Update(dbname *sql.DB, tblname string, params []string) {

	var statementtext = "update "

	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	switch tblname {
	case "Incident": //this is the case for the table of the Incidents
		stmt, err := dbname.Prepare(statementtext + " " + tblname + " set Name= ? , Context= ? , Company= ? , Credential= ? ,Receiving_group = (SELECT Groupid FROM RecieverGroups WHERE Name = ?) , Countermeasure = ? , Sendbymanager=(SELECT Username FROM SystemManager WHERE Username = ?) " + "Where IncidentId=? ;")
		if err != nil {
			fmt.Print("helper_methods.go : 118")
			fmt.Println(err)

		}

		_, queryError := stmt.Exec(params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7])
		if queryError != nil {
			fmt.Print("Something went wrong with the execution of the query")
			fmt.Println(queryError)

		}

	}

}

func Logselect(dbname *sql.DB) {
	var statementtext = "select "
	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	stmt, err := dbname.Prepare(statementtext + " "  + "IncidentId , Name, Context " + " from Incident ;" )
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
		var IncidentId string
		var Name string
		var Context string
		if err := results.Scan(&IncidentId, &Name , &Context); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s %s %s \n", IncidentId ,Name ,Context)

	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Results from select query: ")

}

func Select_warning_receivers(dbname *sql.DB){
	var statementtext = "select "
	//just for proof of concept
	//According to the name of the table we go to the corresponding action and create the appropriate query
	stmt, err := dbname.Prepare(statementtext + " "  + "WriD , Name, PhoneNumber , Company , ReceiverGroup , ReceiverId " + " from WarningReceiver ;" )
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
		var datares[6] string

		if err := results.Scan(&datares[0],&datares[1],&datares[2],&datares[3],&datares[4],&datares[5]); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", datares)

	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Results from select query: ")

}