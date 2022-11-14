package databasefunctions

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	apitools "incidentAPI/apiTools"
	"log"
	"net/http"
)

// generic hashing method without salt if needed
func Hashing(givenText string) string {
	// Here we start with a new hash.
	hasher := sha512.New512_256()
	//conversion of string to byte array for the hasher
	hasher.Write([]byte(givenText))
	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	bs := hasher.Sum(nil)
	//reconstruct the string from the byte array
	newstring := string(bs)
	fmt.Println(newstring)

	return newstring
}

// method used to create salted hashes it is used when we do password check
func Hashingsalted(givenText string, salt string) string {
	salt = "saltOne"

	var passwordBytes = []byte(givenText)
	var saltbytes = []byte(salt)
	// Create sha-512 hasher
	var sha512Hasher = sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, saltbytes...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	// Convert the hashed password to a hex string
	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex

}

// gets the users password from db based on the email he entered and checks it with the password he entered as well after it is being hashed
func Passwdcheck(w http.ResponseWriter, giventext string, email string) int {

	var statementtext = "select "

	results, queryError := Db.Query(statementtext+" "+"Password, CiD from Credentials WHERE Email=?", email)
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Println(queryError.Error())
		return 0
	}

	/*results, queryError := stmt.Query()
	if queryError != nil {
		http.Error(w, apitools.QueryError, http.StatusInternalServerError)
		log.Println(err.Error())
		return 0
	}*/

	defer results.Close()
	fmt.Println("Results from select query: ")
	var Password string
	var CiD string
	for results.Next() {
		if err := results.Scan(&Password, &CiD); err != nil {
			http.Error(w, apitools.DecodeError, http.StatusInternalServerError)
			log.Println(err.Error())
			return 0
		}
	}

	if len(Password) == 0 && len(CiD) == 0 {
		return 0
	}

	if err := results.Err(); err != nil {
		http.Error(w, apitools.UnexpectedError, http.StatusInternalServerError)
		log.Println(err.Error())
		return 0
	}

	result := Hashingsalted(giventext, CiD)

	if result == Password {
		return 1 //if password is correct
	} else {
		return 0 //if password is false
	}
}
