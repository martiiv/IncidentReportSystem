package main

import (
	"fmt"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/endpoints"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

/*
* File main.go starts up the backend, defines enpoint URLs and establishes a databaseconnection
* Uses gorilla mux for routing in order to use variables with front end
? Last rev Martin Iversen 15.11.2022
*/

// Main function initializes router, database and routes URLs
func main() {
	r := mux.NewRouter() //Gorilla mux router

	databasefunctions.EstablishConnection() //Establishes databaseconnection (Uses NTNU MYSQL needs Cisco VPN)

	//Login endpoint for front-end
	r.Path("/login").HandlerFunc(endpoints.LoginSystemManager)

	//ReceiverGroup endpoints
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup).Queries("id", "{id}") //GET
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup)                       //GET, POST, DELETE

	//Incident endpoints
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest).Queries("id", "{id}")   //GET PUT
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest).Queries("tag", "{tag}") //GET PUT
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)                         //POST, DELETE, GET

	//System Manager endpoints
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest).Queries("id", "{id}") //PUT
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest)                       //GET, POST, DELETE

	//Warning Receiver endpoint
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver).Queries("id", "{id}") //PUT, GET
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver)                       //POST, DELETE, GET
	http.Handle("/", r)

	fmt.Print("Listening on port:" + getPort()) //Returns port

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
}

/*
Function used for setting the port for the application
We use localhost 8080 for testing
Takes no parameters
Returns the port the software is listening on
*/
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
