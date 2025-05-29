package presenter

import (
	"encoding/json"
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/rpc/pb"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/module/auth/presenter/dto"
	"github.com/go-chi/render"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"io"
	"net/http"
)

func (p *authPresenter) UserVerifyOtp(w http.ResponseWriter, r *http.Request) {
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	ctx, span := p.telemetryInfrastructure.Tracer(ctx, "Presenter.UserVerifyOtp")
	defer span.End()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error reading body: %v", err))
		render.Status(r, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	req := &dto.UserVerifyOtpRequest{}
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

	resp, err := p.authUseCase.UserVerifyOtp(ctx, r.Header.Get(enum.XRequestIDHeader.String()), &pb.UserVerifyOtpRequest{
		Otp: req.Otp,
	})
	if err != nil {
		p.logger.Error(fmt.Sprintf("error verifying otp: %v", err))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, resp)
	return
}
