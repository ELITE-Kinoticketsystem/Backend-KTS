package utils

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
	"github.com/mailgun/mailgun-go/v4"
	qrsvg "github.com/wamuir/svg-qr-code"
)

type MailgunInterface interface {
	Send(ctx context.Context, message *mailgun.Message) (mes string, id string, err error)
	NewMessage(from, subject, text string, to ...string) *mailgun.Message
}

const welcomeEmailTmpl string = `<!DOCTYPE html>
<html>
<head>
  <title>Welcome to Cinemika</title>
</head>
<body style="background-color: #1A1F25; color: #FFFFFF; font-family: sans-serif;">
  <div class="container" style="margin: 0; width: 100%; background-color: #2A313A; border-radius: 10px;">
    <div class="header" style="color: #FAFAFA; padding: 10px; text-align: center; border-top-left-radius: 10px; border-top-right-radius: 10px;">
      <h1>Welcome to Cinemika</h1>
    </div>
    <div class="content" style="padding: 20px; font-size:15px; color: #FFFFFF;">
      <p style="margin-top:0;color:#b9bec7;font-size:16px;line-height:1.5em">Hi {{ .Username}}, welcome to Cinemika! We're thrilled to have you join our community.</p>
      <!-- Dynamic Content -->
      
      <p>If you have any questions or need assistance, don't hesitate to reach out. Our team is always here to help you make the most of your experience with Cinemika.</p>

      <p><a href="https://cinemika.westeurope.cloudapp.azure.com/auth/login" class="button" style="background-color: #89a3be; color: #FAFAFA; padding: 10px 20px; text-align: center; text-decoration: none; display: inline-block; border-radius: 5px;">Login to Your Account</a></p>
    </div>
  </div>
</body>
</html>`

func PrepareWelcomeMailBody(username string) (string, error) {
	t, err := template.New("welcomeEmail").Parse(welcomeEmailTmpl)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, struct {
		Username string
	}{
		Username: username,
	}); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

const orderConfirmationtmpl string = `<!DOCTYPE html>
<html>
<head>
  <title>Order Confirmation</title>
</head>
<body style="background-color: #1A1F25; color: #FAFAFA; font-family: Arial, sans-serif;">
  <div class="container" style="margin: 0 auto; width: 100%; background-color: #2A313A; padding: 20px; border-radius: 10px;">
    <div class="header" style="color: #FAFAFA; padding: 10px; text-align: center; border-top-left-radius: 10px; border-top-right-radius: 10px;">
      <h1>Order Confirmation</h1>
    </div>
    <div class="content" style="padding: 20px; font-size:15px;">
      <p style="margin-top:0;color:#74787E;font-size:16px;line-height:1.5em">Thank you for booking with CinemaWorld! Your order confirmation is below.</p>
      <!-- Dynamic Content -->
      <p><strong>Event Title:</strong> {{.Event.Title}}</p>
      <p><strong>Cinema Hall:</strong> {{.CinemaHall.Name}}</p>
      <p><strong>Date and Time:</strong> {{prettierTime .Event.Start}}</p>
      <p><strong>Total Price:</strong> {{prettyPrice .Order.Totalprice}}</p>
      <p><strong>Theatre Name:</strong> {{.Theatre.Name}}</p>
      
      <!-- Ticket information table -->
      <h1>Seats</h1>
      <table style="width: 100%; border-collapse: collapse;">
        <tr>
          <th style="padding: 8px; text-align: left; font-size: 13px; border-bottom: 1px solid #74787E;">Price Category</th>
          <th style="padding: 8px; text-align: left;">Seat Category</th>
          <th style="padding: 8px; text-align: left;">Row</th>
          <th style="padding: 8px; text-align: left;">Column</th>
        </tr>
        {{range .Tickets}}
        <tr>
          <td>{{.PriceCategory.CategoryName}}</td>
          <td>{{.SeatCategory.CategoryName}}</td>
          <td>{{.Seat.RowNr}}</td>
          <td>{{.Seat.ColumnNr}}</td>
        </tr>
        {{end}}
      </table>
      <!-- Display QR Code -->
      <h1>QR Code</h1>
      <div class="qr-code-container" style="text-align: center; padding: 10px;">
        {{qrCodeSvgHtmlTemplate .Order.ID}}
      </div>
    </div>
  </div>
</body>
</html>`

func prettierTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

func prettyPrice(price int32) string {
	x := float64(price)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

func qrCodeSvgHtmlTemplate(orderId *uuid.UUID) template.HTML {
	qr, err := qrsvg.New(orderId.String())
	qr.Blocksize = 6
	if err != nil {
		panic(err)
	}

	qr.Borderwidth = 0
	qrSvg := qr.SVG()
	var svgString string = qrSvg.String()

	return template.HTML(svgString)
}

func PrepareOrderConfirmationBody(order models.GetOrderDTO) (string, error) {

	// Parse the template
	t, err := template.New("orderConfirmation").Funcs(template.FuncMap{
		"qrCodeSvgHtmlTemplate": qrCodeSvgHtmlTemplate,
		"prettierTime":          prettierTime,
		"prettyPrice":           prettyPrice,
	}).Parse(orderConfirmationtmpl)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, order); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
