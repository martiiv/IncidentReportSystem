package communication

import (
	"bytes"
	"encoding/json"
	apitools "incidentAPI/apiTools"
	"incidentAPI/structs"
	"log"
	"net/http"
)

/*
* File discord.go
* Uses a discord webhook to send incident reports to warnining receivers
! NB! ONLY PROOF OF CONCEPT AND PARTIAL IMPLEMENTATION
! NOT IMPLEMENTED FULLY DUE TO TIME CONSTRAINTS AND FINNICKY INTEGRATION WITH DISCORD SERVERS
? Last rev Martin Iversen 15.11.2022
*/

// Function defines the discord url and lets the API send incidents via discord
func discordWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	webhookURL := "https://discord.com/api/webhooks/1026402234877087784/8SYbrxxZU0ehZowC6x3qZr-XC890Vi-ddRXwCz8R3H3RKdGN-lYQb4nb9yzMvssZNgif"

	var input structs.MessageInput //Defines the structure for the json body

	err := json.NewDecoder(r.Body).Decode(&input) //Decodes the body into the struct
	if err != nil {
		http.Error(w, apitools.DecodeError, http.StatusBadRequest)
		log.Fatal(err.Error())
		return
	}

	//!Started configuring structure
	var embeds = structs.Embeds{
		Author: structs.Author{
			Name:    "",
			URL:     "",
			IconURL: "",
		},
		Title:       "",
		Description: "",
		Color:       0,
	}

	//!Started configuring structure
	var outputStruct = structs.MessageOutput{
		Username:  "",
		AvatarURL: "",
		Embeds:    embeds,
	}

	var output bytes.Buffer
	err = json.NewEncoder(&output).Encode(outputStruct) //Encodes output
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusInternalServerError)
		log.Fatal(err)
	}

	r, err = http.NewRequest("POST", webhookURL, &output) //Sends the request
	if err != nil {
		http.Error(w, apitools.EncodeError, http.StatusBadRequest)
		panic(err)
	}
}
