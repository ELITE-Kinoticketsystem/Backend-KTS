package schemas

type Order struct {
	ID           int           `json:"id"`
	Total        int           `json:"total"`
	Timestamp    string        `json:"timestamp"` //TODO: find better datatype
	Confirmation *Confirmation `json:"confirmation"`
}

type Confirmation struct {
	Order
	Payment *Payment `json:"payment"`
}

type Reservation struct {
	Order
	// TODO: Payment maybe without reservation?
}

type Cart struct {
	LastEdit string    `json:"lastEdit"`
	Tickets  *[]Ticket `json:"tickets"`
}
