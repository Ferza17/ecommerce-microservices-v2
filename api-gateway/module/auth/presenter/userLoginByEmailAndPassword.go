package presenter

import (
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/presenter/dto"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

func (p *authPresenter) UserLoginByEmailAndPassword(w http.ResponseWriter, r *http.Request) {
	ctx, span := p.telemetryInfrastructure.Tracer(r.Context(), "Presenter.UserLoginByEmailAndPasswordRequest")
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

	if _, err = p.authUseCase.UserLoginByEmailAndPassword(ctx, r.Header.Get(enum.XRequestIDHeader.String()), &pb.UserLoginByEmailAndPasswordRequest{
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
