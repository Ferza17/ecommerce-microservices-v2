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

func (p *authPresenter) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.CreateUser")
	defer span.End()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error reading body: %v", err))
		render.Status(r, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	req := &dto.CreateUserRequest{}
	res := &dto.Response{}
	if err = json.Unmarshal(body, req); err != nil {
		p.logger.Error(fmt.Sprintf("error unmarshaling body: %v", err))
		if err = res.WriteResponse(w, http.StatusInternalServerError, "Internal Server Error"); err != nil {
			p.logger.Error(fmt.Sprintf("Error while creating response %v", err))
		}
		return
	}
	if err = req.Validate(); err != nil {
		p.logger.Error(fmt.Sprintf("error validating request: %v", err))
		if err = res.WriteResponse(w, http.StatusBadRequest, err.Error()); err != nil {
			p.logger.Error(fmt.Sprintf("Error while creating response %v", err))
		}
		return
	}

	requestId := r.Header.Get(enum.XRequestIDHeader.String())
	if _, err = p.userUseCase.CreateUser(ctx, requestId, &userRpc.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}); err != nil {
		p.logger.Error(fmt.Sprintf("error creating user: %v", err))
		if err = res.WriteResponse(w, http.StatusInternalServerError, "Internal Server Error"); err != nil {
			p.logger.Error(fmt.Sprintf("Error while creating response %v", err))
		}
		return
	}

	if err = res.WriteResponse(w, http.StatusCreated, "success"); err != nil {
		p.logger.Error(fmt.Sprintf("Error while creating response %v", err))
	}
	return
}
