package structs

type CreateWarningReceiver struct {
	Name          string `json:"name"`
	PhoneNumber   string `json:"phoneNumber"`
	Company       string `json:"company"`
	ReceiverGroup string `json:"receiverGroup"`
	ReceiverEmail string `json:"receiverEmail"`
}

type GetWarningReceiver struct {
	Id            int     `json:"id"`
	Name          *string `json:"name"`
	PhoneNumber   *string `json:"phoneNumber"`
	Company       *string `json:"company"`
	ReceiverGroup *string `json:"receiverGroup"`
	ReceiverEmail *string `json:"receiverEmail"`
}

type DeleteWarningReceiver []struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type GetEmails struct {
	Email string `json:"Email"`
}
