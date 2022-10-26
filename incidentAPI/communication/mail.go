package communication

import (
	"encoding/json"
	"fmt"
	_ "incidentAPI/apiTools"
	"incidentAPI/config"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
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
	//receiver := incident.Receiver

	to := getEmails(incident.Receiver)

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

	var project structs.MessageInput        //Defines the structure of the request
	err = json.Unmarshal(request, &project) //Unmarshall the request data into the project struct
	if err != nil {
		fmt.Println("Error")
	}

	return project
}

func getEmails(groups []int) []string {
	var emails []string

	var searchString string
	if len(groups) > 1 {
		for i := 0; i < len(groups); i++ {
			searchString += strconv.Itoa(groups[i]) + " AND "
		}
	}

	searchString = searchString[:len(searchString)-5]

	fmt.Println(searchString)

	row, err := databasefunctions.Db.Query("SELECT Emails.Email FROM WarningReceiver INNER JOIN Emails ON WarningReceiver.ReceiverEmail = Emails.Email INNER JOIN ReceiverGroups ON WarningReceiver.ReceiverGroup = ReceiverGroups.Name WHERE ReceiverGroups.Groupid = ?", searchString)
	if err != nil {
		return nil
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for row.Next() {
		email := structs.GetEmails{}
		if err = row.Scan(
			&email.Email,
		); err != nil {
			return nil
		}

		emails = append(emails, email.Email)

	}
	return emails

}
