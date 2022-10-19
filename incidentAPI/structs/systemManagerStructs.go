package structs

type CreateWarningReceiver struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Company     string `json:"company"`
	Group       string `json:"group"`
}

type GetWarningReceiver struct {
	Id           int     `json:"id"`
	Name         *string `json:"name"`
	PhoneNumber  *string `json:"phoneNumber"`
	Company      *string `json:"company"`
	CredentialId *string `json:"credentialId"`
	Group        *string `json:"group"`
	ReceiverId   *string `json:"receiverId"`
	Email        *string `json:"email"`
}

type DeleteWarningReceiver []struct {
	Id int `json:"id"`
}
