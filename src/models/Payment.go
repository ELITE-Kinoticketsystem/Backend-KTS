package main

type Payment struct {
	PaymentMethod *PaymentMethod `json:"paymentMethod"`
	Invoice       Invoice        `json:"Invoice"`
}

type PaymentMethod struct {
	// TODO: information
}

type Mastercard struct {
	PaymentMethod
}

type PayPal struct {
	PaymentMethod
}

type ApplePay struct {
	PaymentMethod
}

type Visa struct {
	PaymentMethod
}

type Cash struct {
	PaymentMethod
}

type Invoice struct {
	// TODO:
}