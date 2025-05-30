package presenter

import (
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/gen/user/v1"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/presenter/dto"
	"github.com/go-chi/render"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
)

func (p *authPresenter) UserLoginByEmailAndPassword(w http.ResponseWriter, r *http.Request) {
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.UserLoginByEmailAndPassword")
	defer span.End()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error reading body: %v", err))
		render.Status(r, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	req := &dto.UserLoginByEmailAndPasswordRequest{}
	if err = json.Unmarshal(body, req); err != nil {
		p.logger.Error(fmt.Sprintf("error unmarshaling body: %v", err))
		render.Status(r, http.StatusInternalServerError)
		return
	}
	if err = req.Validate(); err != nil {
		p.logger.Error(fmt.Sprintf("error validating request: %v", err))
		render.Status(r, http.StatusBadRequest)
		return
	}

	if err = p.authUseCase.UserLoginByEmailAndPassword(ctx, r.Header.Get(enum.XRequestIDHeader.String()), &userRpc.UserLoginByEmailAndPasswordRequest{
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		p.logger.Error(fmt.Sprintf("error creating user: %v", err))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	return
}
