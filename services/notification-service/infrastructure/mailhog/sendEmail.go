package mailhog

import (
	"fmt"
	"net/smtp"

	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/util"
	"github.com/jordan-wright/email"
)

func (i *mailhogInfrastructure) SendMail(mailer *Mailer) error {

	templateBuffer, err := util.ParseEmailTemplate(mailer.Template, mailer.TemplateVars)
	if err != nil {
		i.logger.Error(fmt.Sprintf("Failed to parse email template: %v", err))
		return err
	}

	emailClient := email.NewEmail()
	emailClient.From = config.Get().ConfigSmtp.SenderEmail
	emailClient.Bcc = []string{config.Get().ConfigSmtp.SenderEmail}
	emailClient.Cc = []string{config.Get().ConfigSmtp.SenderEmail}
	emailClient.To = []string{mailer.To}
	emailClient.Subject = mailer.Subject
	emailClient.HTML = templateBuffer.Bytes()

	if err = emailClient.Send(
		fmt.Sprintf("%s:%s", config.Get().ConfigSmtp.Host, config.Get().ConfigSmtp.Port),
		smtp.CRAMMD5Auth(config.Get().ConfigSmtp.Username, config.Get().ConfigSmtp.Password),
	); err != nil {
		i.logger.Error(fmt.Sprintf("Failed to send email: %v", err))
		return err
	}

	return nil
}
