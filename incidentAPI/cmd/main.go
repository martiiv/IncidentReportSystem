package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	establishConnection()
	r := mux.NewRouter()

	//Log endpoint
	r.Path("incident")                 //GET, POST
	r.Path("incident/countermeasures") //PUT

	//System Manager endpoint
	r.Path("manager") //GET, POST, DELETE

	//Warning Receiver endpoint
	r.Path("warningreveiver") //POST, DELETE

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		fmt.Printf(err.Error())
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
