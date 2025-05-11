package mailhog

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func (i *mailhogInfrastructure) SendMail(mailer *Mailer) error {

	templateBuffer, err := util.ParseEmailTemplate(mailer.Template, mailer.TemplateVars)
	if err != nil {
		i.logger.Error(fmt.Sprintf("Failed to parse email template: %v", err))
		return err
	}

	emailClient := email.NewEmail()
	emailClient.From = config.Get().SmtpSenderEmail
	emailClient.Bcc = []string{config.Get().SmtpSenderEmail}
	emailClient.Cc = []string{config.Get().SmtpSenderEmail}
	emailClient.To = []string{mailer.To}
	emailClient.Subject = mailer.Subject
	emailClient.HTML = templateBuffer.Bytes()

	if err := emailClient.Send(
		fmt.Sprintf("%s:%d", config.Get().SmtpHost, config.Get().SmtpPort),
		smtp.CRAMMD5Auth(config.Get().SmtpUsername, config.Get().SmtpPassword),
	); err != nil {
		i.logger.Error(fmt.Sprintf("Failed to send email: %v", err))
		return err
	}

	return nil
}
