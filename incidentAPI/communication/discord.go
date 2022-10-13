package communication

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func discordWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	webhookURL := "https://discord.com/api/webhooks/1026402234877087784/8SYbrxxZU0ehZowC6x3qZr-XC890Vi-ddRXwCz8R3H3RKdGN-lYQb4nb9yzMvssZNgif"

	var input messageInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var embeds = Embeds{
		Author: Author{
			Name:    "",
			URL:     "",
			IconURL: "",
		},
		Title:       "",
		Description: "",
		Color:       0,
	}

	var outputStruct = messageOutput{
		Username:  "",
		AvatarURL: "",
		Embeds:    embeds,
	}

	var output bytes.Buffer
	err = json.NewEncoder(&output).Encode(outputStruct)
	if err != nil {
		log.Fatal(err)
	}

	r, err = http.NewRequest("POST", webhookURL, &output)
	if err != nil {
		panic(err)
	}
}
