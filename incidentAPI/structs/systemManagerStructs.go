package structs

type CreateWarningReciever struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Company     string `json:"company"`
	Group       string `json:"group"`
}

type GetWarningReciever struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Company     string `json:"company"`
	Group       string `json:"group"`
}

type DeleteWarningReceiever []struct {
	Id int `json:"id"`
}
