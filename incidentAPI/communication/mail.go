package communication

import (
	"fmt"
	_ "incidentAPI/apiTools"
	"incidentAPI/config"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/structs"
	_ "io/ioutil"
	"log"
	"net/smtp"
)

/*
*
Function that will send mail
*/
func SendMail(inputStruct structs.CreateIncident) {
	// Sending data from email
	from := "trakkemaskintrine@gmail.com"
	password := config.SenderEmailAppPassword

	//Array of email receivers
	to := getEmails(inputStruct.ReceivingGroup)

	//smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	//Main mail body
	messageString := "Subject: " + inputStruct.Tag + "\n\nInformation:\n" + inputStruct.Description + "\r\nCountermeasures:\n" + inputStruct.Countermeasure

	// Message.
	message := []byte((messageString))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println("Mail sent")

}

/*
*
Function that will retrieve the emails of a selected group
*/
func getEmails(groups string) []string {
	var emails []string

	//Create SQL Query statement
	var searchString string
	/*	if len(groups) > 1 {
		for i := 0; i < len(groups); i++ {
			searchString += strconv.Itoa(groups[i]) + " AND "
		}
	}*/

	searchString = groups
	//searchString = searchString[:len(searchString)-5] //Will trim the last AND

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
