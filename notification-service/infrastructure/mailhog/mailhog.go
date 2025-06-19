package mailhog

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
)

type (
	IMailhogInfrastructure interface {
		SendMail(mailer *Mailer) error
	}

	Mailer struct {
		Subject      string
		To           string
		Template     string
		TemplateVars map[string]any
	}
	mailhogInfrastructure struct {
		logger pkg.IZapLogger
	}
)

func NewMailhogInfrastructure(logger pkg.IZapLogger) IMailhogInfrastructure {
	return &mailhogInfrastructure{
		logger: logger,
	}
}
