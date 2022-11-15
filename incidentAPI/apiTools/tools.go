package apitools

import (
	"fmt"
	"strconv"
	"strings"
)

/*
* File tools.go
* Created for generic helper functions used for all Back end code
? Last rev Martin Iversen 15.11.2022
*/

// Functioon correctQuery used for checking query structure before querying the database
func CorrectQuery(query string) (bool, string) {
	queryArr := strings.Split(query, "=") //Splits query into an array
	id := queryArr[1]

	_, err := strconv.Atoi(id) //Converts id to int

	if queryArr[0] == "id" && err == nil { //Logic for checking queryArray
		fmt.Println(queryArr)
	} else {
		fmt.Println("False")
		return false, ""
	}

	return true, id
}
