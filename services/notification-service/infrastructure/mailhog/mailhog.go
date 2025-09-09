package mailhog

import (
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/pkg/logger"
	"github.com/google/wire"
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
		logger logger.IZapLogger
	}
)

var Set = wire.NewSet(NewMailhogInfrastructure)

func NewMailhogInfrastructure(logger logger.IZapLogger) IMailhogInfrastructure {
	return &mailhogInfrastructure{
		logger: logger,
	}
}
