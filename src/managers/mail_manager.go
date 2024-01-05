package managers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type MailMgr interface {
}

type MailManager struct {
	MailgunInstance *mailgun.MailgunImpl
}

const retryMailCount = 3
const emailSender = "Cinemika Team <team@cinemika.tech>"
const emailReceiver = "team@cinemika.tech"

func (mm *MailManager) SendWelcomeMail(to []string) error {
	subject := "Welcome to Cinemika!"
	body := "Welcome to Cinemika! We are glad to have you here."

	return mm.sendMail(to, emailSender, subject, body)
}

func (mm *MailManager) sendMail(to []string, from, subject, body string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	message := mm.MailgunInstance.NewMessage(emailSender, subject, "", to...)
	message.AddHeader("Content-Type", "text/html")
	message.SetHtml(body)

	_, _, err := mm.MailgunInstance.Send(ctx, message)
	if err != nil {
		log.Println("Error in MailManager.SendMail().MailgunInstance.Send(): ", err.Error())
		return err
	}

	return nil
}

func InitializeMailgunClient() *mailgun.MailgunImpl {
	ApiKey := os.Getenv("MAILGUN_API_KEY")
	Domain := os.Getenv("MAILGUN_DOMAIN")

	log.Println("Initializing Mailgun client...")
	log.Println("Domain: ", Domain)

	mg := mailgun.NewMailgun(Domain, ApiKey)
	mg.SetAPIBase(mailgun.APIBaseEU)

	return mg
}
