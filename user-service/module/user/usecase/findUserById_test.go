package usecase

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	postgresInfra "github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/rabbitmq"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/infrastructure/telemetry"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/orm"
	userRpc "github.com/ferza17/ecommerce-microservices-v2/user-service/model/rpc/gen/v1/user"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/auth/repository/redis"
	rolePgRepo "github.com/ferza17/ecommerce-microservices-v2/user-service/module/role/repository/postgres"
	userPgRepo "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgres"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/pkg/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func Test_userUseCase_FindUserById(t *testing.T) {

	mockTelemetryInfrastructure := telemetry.NewITelemetryInfrastructureMock(t)
	mockPostgresSQLInfrastructure := postgresInfra.NewIPostgresSQLMock(t)
	mockUserPostgresqlRepository := userPgRepo.NewIUserPostgresqlRepositoryMock(t)
	mockLogger := logger.NewIZapLoggerMock(t)

	type fields struct {
		userPostgresqlRepository  userPgRepo.IUserPostgresqlRepository
		rolePostgresqlRepository  rolePgRepo.IRolePostgresqlRepository
		rabbitmqInfrastructure    rabbitmq.IRabbitMQInfrastructure
		postgresSQLInfrastructure postgresInfra.IPostgresSQL
		telemetryInfrastructure   telemetry.ITelemetryInfrastructure
		authRedisRepository       redis.IAuthRedisRepository
		logger                    logger.IZapLogger
	}
	type args struct {
		ctx       context.Context
		requestId string
		req       *userRpc.FindUserByIdRequest
	}
	type Err struct {
		ErrThisFunc     error
		ErrFindUserById error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *userRpc.FindUserByIdResponse
		wantErr Err
	}{
		{
			name: "Success",
			fields: fields{
				userPostgresqlRepository:  mockUserPostgresqlRepository,
				rolePostgresqlRepository:  rolePgRepo.NewIRolePostgresqlRepositoryMock(t),
				rabbitmqInfrastructure:    rabbitmq.NewIRabbitMqInfrastructureMock(t),
				postgresSQLInfrastructure: mockPostgresSQLInfrastructure,
				telemetryInfrastructure:   mockTelemetryInfrastructure,
				authRedisRepository:       redis.NewIAuthRedisRepositoryMock(t),
				logger:                    logger.NewIZapLoggerMock(t),
			},
			args: args{
				ctx:       context.Background(),
				requestId: "1234",
				req: &userRpc.FindUserByIdRequest{
					Id: "abc",
				},
			},
			want: &userRpc.FindUserByIdResponse{
				Error:   "",
				Message: "abc",
				Code:    uint32(codes.OK),
				Data: &userRpc.FindUserByIdResponse_FindUserByIdResponseData{
					User: &userRpc.User{Id: "abc"},
				},
			},
			wantErr: Err{
				ErrThisFunc:     nil,
				ErrFindUserById: nil,
			},
		},
		{
			name: "Error FindUserById ErrRecordNotFound",
			fields: fields{
				userPostgresqlRepository:  mockUserPostgresqlRepository,
				rolePostgresqlRepository:  rolePgRepo.NewIRolePostgresqlRepositoryMock(t),
				rabbitmqInfrastructure:    rabbitmq.NewIRabbitMqInfrastructureMock(t),
				postgresSQLInfrastructure: mockPostgresSQLInfrastructure,
				telemetryInfrastructure:   mockTelemetryInfrastructure,
				authRedisRepository:       redis.NewIAuthRedisRepositoryMock(t),
				logger:                    mockLogger,
			},
			args: args{
				ctx:       context.Background(),
				requestId: "1234",
				req: &userRpc.FindUserByIdRequest{
					Id: "",
				},
			},
			want: &userRpc.FindUserByIdResponse{
				Error:   gorm.ErrRecordNotFound.Error(),
				Message: "abc",
				Code:    uint32(codes.OK),
				Data: &userRpc.FindUserByIdResponse_FindUserByIdResponseData{
					User: &userRpc.User{Id: "abc"},
				},
			},
			wantErr: Err{
				ErrThisFunc:     status.Error(codes.NotFound, "User not found"),
				ErrFindUserById: gorm.ErrRecordNotFound,
			},
		},
		{
			name: "Error FindUserById ErrInternalServerError",
			fields: fields{
				userPostgresqlRepository:  mockUserPostgresqlRepository,
				rolePostgresqlRepository:  rolePgRepo.NewIRolePostgresqlRepositoryMock(t),
				rabbitmqInfrastructure:    rabbitmq.NewIRabbitMqInfrastructureMock(t),
				postgresSQLInfrastructure: mockPostgresSQLInfrastructure,
				telemetryInfrastructure:   mockTelemetryInfrastructure,
				authRedisRepository:       redis.NewIAuthRedisRepositoryMock(t),
				logger:                    mockLogger,
			},
			args: args{
				ctx:       context.Background(),
				requestId: "1234",
				req: &userRpc.FindUserByIdRequest{
					Id: "",
				},
			},
			want: &userRpc.FindUserByIdResponse{
				Error:   gorm.ErrRecordNotFound.Error(),
				Message: "abc",
				Code:    uint32(codes.OK),
				Data: &userRpc.FindUserByIdResponse_FindUserByIdResponseData{
					User: &userRpc.User{Id: "abc"},
				},
			},
			wantErr: Err{
				ErrThisFunc:     status.Error(codes.Internal, "error"),
				ErrFindUserById: gorm.ErrPreloadNotAllowed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUseCase{
				userPostgresqlRepository:  tt.fields.userPostgresqlRepository,
				rolePostgresqlRepository:  tt.fields.rolePostgresqlRepository,
				rabbitmqInfrastructure:    tt.fields.rabbitmqInfrastructure,
				postgresSQLInfrastructure: tt.fields.postgresSQLInfrastructure,
				telemetryInfrastructure:   tt.fields.telemetryInfrastructure,
				authRedisRepository:       tt.fields.authRedisRepository,
				logger:                    tt.fields.logger,
			}

			mockSpan := trace.SpanFromContext(tt.args.ctx)
			mockTelemetryInfrastructure.
				OnStartSpanFromContext(mock.Anything).
				TypedReturns(tt.args.ctx, mockSpan)

			mockDb, _, _ := sqlmock.New()
			dialector := postgres.New(postgres.Config{
				Conn:       mockDb,
				DriverName: "postgres",
			})
			db, _ := gorm.Open(dialector, &gorm.Config{})
			mockPostgresSQLInfrastructure.
				OnGormDB().
				TypedReturns(db)

			mockUserPostgresqlRepository.
				On("FindUserById", tt.args.requestId, tt.args.req.Id, mock.AnythingOfType("*gorm.DB")).
				Return(&orm.User{ID: tt.args.req.Id}, tt.wantErr.ErrThisFunc)

			if tt.want != nil {
				mockLogger.
					On("Error", mock.Anything, mock.Anything, mock.Anything).
					Return()
			}

			_, err := u.FindUserById(tt.args.ctx, tt.args.requestId, tt.args.req)
			if tt.wantErr.ErrThisFunc == nil {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
