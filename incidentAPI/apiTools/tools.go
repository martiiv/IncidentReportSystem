package apitools

import (
	"fmt"
	"strconv"
	"strings"
)

func CorrectQuery(query string) (bool, string) {
	queryArr := strings.Split(query, "=")
	id := queryArr[1]

	_, err := strconv.Atoi(id)

	if queryArr[0] == "id" && err == nil {
		fmt.Println(queryArr)
	} else {
		fmt.Println("False")
		return false, ""
	}

	return true, id
}
