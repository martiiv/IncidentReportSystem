package communication

import (
	"fmt"
	apitools "incidentAPI/apiTools"
	"incidentAPI/config"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	_ "io"
	"log"
	"net/http"
	"net/smtp"
)

/*
* File mail.go
* Lets a system manager send mails with information regarding incidents at the company
* NB Predefined emails used for testing only if deployed another mail needs to be used
? Last revision Martin Iversen 15.11.2022
*/

// Function sendmail sends an email with information regarding an incident
func SendMail(w http.ResponseWriter, inputStruct structs.CreateIncident) error {
	// Sending data from email
	from := "trakkemaskintrine@gmail.com" //! Predefined mail needs to be changed if system is deployed
	password := config.SenderEmailAppPassword

	//Array of email receivers
	to := getEmails(inputStruct.ReceivingGroup)

	//smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Main mail body
	messageString := "Subject: " + inputStruct.Tag + "\n\nInformation:\n" + inputStruct.Description + "\n\nCountermeasure:\n" + inputStruct.Countermeasure

	// Message.
	message := []byte((messageString))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
		log.Println(err.Error())
		return err
	}

	fmt.Fprint(w, "Mail sent")
	return err
}

/*
* Function that will retrieve the emails of a selected group
 */
func getEmails(groups string) []string {
	var emails []string

	//Create SQL Query statement
	searchString := groups

	//Fetching from database
	row, err := databasefunctions.Db.Query("SELECT Emails.Email FROM WarningReceiver INNER JOIN Emails ON WarningReceiver.ReceiverEmail = Emails.Email INNER JOIN ReceiverGroups ON WarningReceiver.ReceiverGroup = ReceiverGroups.Name WHERE ReceiverGroups.Name = ?", searchString)
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
