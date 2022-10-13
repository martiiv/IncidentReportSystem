package communication

import (
	"encoding/json"
	"fmt"
	"incidentAPI/config"
	"incidentAPI/structs"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
)

func SendMail(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Du er inne i sendmail")

	var incident structs.SendIndividualIncident
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&incident)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Du har decoda")

	// Sender data.
	from := "trakkemaskintrine@gmail.com"
	password := config.SenderEmailAppPassword

	// Receiver email address.
	reciever := incident.Receiver

	to := []string{reciever}

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
