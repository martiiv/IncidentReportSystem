package communication

import (
	"encoding/json"
	_ "incidentAPI/apiTools"
	"incidentAPI/config"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	_ "io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
)

/**
Function that will send mail
*/
func SendMail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	//Declaring struct to
	var incident structs.SendIndividualIncident
	err := json.NewDecoder(r.Body).Decode(&incident) //Decodes the requests body into the structure defined above
	if err != nil {
		http.Error(w, "email not sent", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	// Sending data from email
	from := "trakkemaskintrine@gmail.com"
	password := config.SenderEmailAppPassword

	//Array of email receivers
	to := getEmails(incident.Receiver)

	//smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Main mail body
	messageString := "Subject: " + incident.Context + "\n\nInformation:\n" + incident.Information + "\r\nCountermeasures:\n" + incident.Countermeasure

	// Message.
	message := []byte((messageString))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		http.Error(w, "email not sent", http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	http.Error(w, "email successfully sent", http.StatusOK)
}

/**
Function that will retrieve the emails of a selected group
*/
func getEmails(groups []int) []string {
	var emails []string

	//Create SQL Query statement
	var searchString string
	if len(groups) > 1 {
		for i := 0; i < len(groups); i++ {
			searchString += strconv.Itoa(groups[i]) + " AND "
		}
	}
	searchString = searchString[:len(searchString)-5] //Will trim the last AND

	//Fetching from database
	row, err := databasefunctions.Db.Query("SELECT Emails.Email FROM WarningReceiver INNER JOIN Emails ON WarningReceiver.ReceiverEmail = Emails.Email INNER JOIN ReceiverGroups ON WarningReceiver.ReceiverGroup = ReceiverGroups.Name WHERE ReceiverGroups.Groupid = ?", searchString)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for row.Next() {
		email := structs.GetEmails{}
		if err = row.Scan(
			&email.Email,
		); err != nil {
			log.Println(err.Error())
			return nil
		}

		emails = append(emails, email.Email)

	}
	return emails

}
