package workflow

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/temporal"
	pb "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/usecase"
	"github.com/google/wire"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/emptypb"
)

type (
	IAuthWorkflow interface {
		AuthUserRegisterWorkflow(wCtx workflow.Context, requestId string, req *pb.AuthUserRegisterRequest) (*emptypb.Empty, error)
	}
	authWorkflow struct {
		temporal    temporal.ITemporalInfrastructure
		authUseCase usecase.IAuthUseCase
	}
)

var Set = wire.NewSet(NewAuthWorkflow)

func NewAuthWorkflow(
	temporal temporal.ITemporalInfrastructure,
	authUseCase usecase.IAuthUseCase,
) IAuthWorkflow {
	c := &authWorkflow{
		temporal:    temporal,
		authUseCase: authUseCase,
	}
	c.temporal = c.temporal.RegisterWorkflow(c.AuthUserRegisterWorkflow)
	return c
}
