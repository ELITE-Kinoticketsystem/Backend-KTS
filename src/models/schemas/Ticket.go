package schemas

type Ticket struct {
	ID            int           `json:"id"`
	Price         float64       `json:"price"`
	Timestamp     string        `json:"timestamp"` //TODO: find better datatype
	Validated     bool          `json:"validated"`
	Paid          bool          `json:"paid"`
	Reserved      bool          `json:"reserved"`
	Seat          *Seat         `json:"seat"`
	QrCode        string        `json:"qrCode"` // TODO: make maybe as Object, depending on how this works
	PriceCategory PriceCategory `json:"priceCategory"`
	Order         *Order        `json:"order"`
}

type PriceCategory string

const (
	StudentDiscount PriceCategory = "StudentDiscount"
	ChildDiscount   PriceCategory = "ChildDiscount"
	ElderlyDiscount PriceCategory = "ElderlyDiscount"
	RegularPrice    PriceCategory = "RegularPrice"
)