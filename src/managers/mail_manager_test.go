package managers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/mailgun/mailgun-go/v4"
	"go.uber.org/mock/gomock"
)

func TestSendWelcomeMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mg := mailgun.NewMailgun("", "")
	message := mg.NewMessage(emailSender, "Welcome to Cinemika!", "", "test")

	mockMailgungInterface := mocks.NewMockMailgunInterface(mockCtrl)
	mockMailgungInterface.EXPECT().NewMessage(emailSender, "Welcome to Cinemika!", "", "test").Return(message)
	mockMailgungInterface.EXPECT().Send(gomock.Any(), gomock.Any()).Return("", "", nil)

	mailManager := MailManager{
		MailgunInstance: mockMailgungInterface,
	}

	mailManager.SendWelcomeMail("test", "test")

}

func TestSendOrderConfirmationMail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mg := mailgun.NewMailgun("", "")
	message := mg.NewMessage(emailSender, "Order confirmation", "", "test")

	mockMailgungInterface := mocks.NewMockMailgunInterface(mockCtrl)
	mockMailgungInterface.EXPECT().NewMessage(emailSender, "Order confirmation", "", "test").Return(message)
	mockMailgungInterface.EXPECT().Send(gomock.Any(), gomock.Any()).Return("", "", nil)

	mailManager := MailManager{
		MailgunInstance: mockMailgungInterface,
	}

	order := samples.GetOrderSample()

	mailManager.SendOrderConfirmationMail("test", order)

}
