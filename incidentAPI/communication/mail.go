package communication

import (
	"encoding/json"
	"fmt"
	"incidentAPI/config"
	"incidentAPI/structs"
	"io/ioutil"
	"net/http"
	"net/smtp"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	fmt.Print("Du er inne i sendmail")

	var incident structs.SendIndividualIncident
	err := json.NewDecoder(r.Body).Decode(&incident) //Decodes the requests body into the structure defined above
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Print("Du har decoda")

	// Sender data.
	from := "trakkemaskintrine@gmail.com"
	password := config.SenderEmailAppPassword

	// Receiver email address.
	receiver := incident.Receiver

	to := []string{receiver}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	messageString := incident.Information

	// Message.
	message := []byte((messageString))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Email Sent Successfully!")
}

func addStruct(r *http.Request) interface{} {

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error")
	}

	var project messageInput                //Defines the structure of the request
	err = json.Unmarshal(request, &project) //Unmarshall the request data into the project struct
	if err != nil {
		fmt.Println("Error")
	}

	return project
}
