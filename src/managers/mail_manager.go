package managers

import (
	"context"
	"log"
	"os"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/mailgun/mailgun-go/v4"
)

type MailMgr interface {
	SendWelcomeMail(to string, username string) *models.KTSError
	SendOrderConfirmationMail(to string, order models.GetOrderDTO) *models.KTSError
}

type MailManager struct {
	MailgunInstance utils.MailgunInterface
}

const retryMailCount = 3
const emailSender = "Cinemika Team <team@cinemika.tech>"
const emailReceiver = "team@cinemika.tech"

func (mm *MailManager) SendWelcomeMail(to string, username string) *models.KTSError {
	subject := "Welcome to Cinemika!"

	body, err := utils.PrepareWelcomeMailBody(username)
	log.Println(body)
	if err != nil {
		return kts_errors.KTS_UPSTREAM_ERROR
	}

	return mm.sendMail(to, emailSender, subject, body)
}

func (mm *MailManager) SendOrderConfirmationMail(to string, order models.GetOrderDTO) *models.KTSError {
	subject := "Order confirmation"

	body, err := utils.PrepareOrderConfirmationBody(order)
	if err != nil {
		return kts_errors.KTS_UPSTREAM_ERROR
	}

	return mm.sendMail(to, emailSender, subject, body)
}

func (mm *MailManager) sendMail(to string, from, subject, body string) *models.KTSError {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	message := mm.MailgunInstance.NewMessage(emailSender, subject, "", to)
	message.AddHeader("Content-Type", "text/html")
	message.SetHtml(body)

	_, _, err := mm.MailgunInstance.Send(ctx, message)
	if err != nil {
		log.Println("Error in MailManager.SendMail().MailgunInstance.Send(): ", err.Error())
		return kts_errors.KTS_UPSTREAM_ERROR
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
