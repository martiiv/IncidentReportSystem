package endpoints

import "net/http"

func HandleLogRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	method := r.Method

	switch method {
	case "GET":
		returnAccosiatedCountermeasure(w, r)

	case "POST":
		createConnectionCountermeasure(w, r)

	case "PUT":

	case "DELETE":
	}
}

func returnAccosiatedCountermeasure(w http.ResponseWriter, r *http.Request) {

}

func createConnectionCountermeasure(w http.ResponseWriter, r *http.Request) {

}
