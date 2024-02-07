package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/matcornic/hermes/v2"
)

type MailgunInterface interface {
	Send(ctx context.Context, message *mailgun.Message) (mes string, id string, err error)
	NewMessage(from, subject, text string, to ...string) *mailgun.Message
}

var h = hermes.Hermes{
	Product: hermes.Product{
		Name:        "Cinemika",
		Link:        URL,
		TroubleText: "If the {ACTION}-button is not working for you, just copy and paste the URL below into your web browser.",
		Copyright:   "Copyright © 2024 Cinemika-WWI22SEB",
	},
	Theme: new(hermes.Default),
}

func PrepareWelcomeMailBody(username string) (string, error) {
	hermesMail := hermes.Email{
		Body: hermes.Body{
			Name: username,
			Intros: []string{
				fmt.Sprintf("Welcome to Cinemika, %v! We're very excited to have you on board.", username),
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}

func prettierTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func prettyPrice(price int32) string {
	x := float64(price)
	x = x / 100
	return fmt.Sprintf("€%.2f", x)
}

func generateDataFromOrder(order models.GetOrderDTO) [][]hermes.Entry {
	var data [][]hermes.Entry
	for _, ticket := range order.Tickets {
		data = append(data, []hermes.Entry{
			{Key: "Price Category", Value: ticket.PriceCategory.CategoryName},
			{Key: "Seat Category", Value: ticket.SeatCategory.CategoryName},
			{Key: "Row", Value: fmt.Sprint(ticket.Seat.RowNr)},
			{Key: "Column", Value: fmt.Sprint(ticket.Seat.ColumnNr)},
		})
	}
	return data
}

func PrepareOrderConfirmationBody(order models.GetOrderDTO) (string, error) {

	hermesMail := hermes.Email{
		Body: hermes.Body{
			Intros: []string{
				"Your order has been processed successfully.",
			},
			Dictionary: []hermes.Entry{
				{Key: "Event Title", Value: order.Event.Title},
				{Key: "Cinema Hall", Value: order.CinemaHall.Name},
				{Key: "Date and Time", Value: prettierTime(order.Event.Start)},
				{Key: "Total Price", Value: prettyPrice(order.Order.Totalprice)},
				{Key: "Theatre Name", Value: order.Theatre.Name},
			},
			Table: hermes.Table{

				Data: generateDataFromOrder(order),

				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Price Category": "35%",
						"Seat Category":  "35%",
						"Row":            "15%",
						"Column":         "15%",
					},
				},
			},
			Actions: []hermes.Action{
				{
					Instructions: "You can check your order and more in your dashboard:",
					Button: hermes.Button{
						Text:  "Go to Dashboard",
						Link:  URL + "/dashboard",
						Color: "#334155",
					},
				},
			},
		},
	}

	body, err := h.GenerateHTML(hermesMail)

	if err != nil {
		return "", err
	}

	return body, nil
}
