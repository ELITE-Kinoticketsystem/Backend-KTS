package schemas

type Address struct {
	ID       int    `json:"id"`
	Street   string `json:"street"`
	Streetnr string `json:"streetnr"`
	Zipcode  string `json:"zipcode"`
	City     string `json:"city"`
	Country  string `json:"country"`
}