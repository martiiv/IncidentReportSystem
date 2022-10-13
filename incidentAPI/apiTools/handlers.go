package apitools

/*
Class handlers.go
Contains methods for basic http methods
GET, POST, PUT and DELETE
author Martin Iversen
rev date 13.10.2022
*/
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*
Function MasterHandler, method will return the requests body as a string and a byte.
Accepts the method strings GET , POST , PUT and DELETE
*/
func MasterHandler(w http.ResponseWriter, r *http.Request, url string, Method string) (string, []byte) {

	switch r.Method {

	case "GET":
		response, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			log.Print("Bad request URL")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		return string(body), body

	case "POST":
		response, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			log.Fatal(err)
			log.Print("Bad request URL")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		return string(body), body

	case "PUT":
		response, err := http.NewRequest(http.MethodPut, url, nil)
		if err != nil {
			log.Fatal(err)
			log.Print("Bad request URL")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		return string(body), body

	case "DELETE":
		response, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			log.Fatal(err)
			log.Print("Bad request URL")
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		return string(body), body
	}

	return "Incorrect method string used", nil
}

func DecodeBody(w http.ResponseWriter, r *http.Request, structBody struct{}) struct{} {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&structBody)
	if err != nil {
		log.Fatal(err)
	}

	return structBody
}

func EncodeBody(w http.ResponseWriter, r *http.Request, body struct{}) http.ResponseWriter {
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Fatal(err)
	}

	return w
}
