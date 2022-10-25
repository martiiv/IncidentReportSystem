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
	Id int `json:"id"`
}
