package grpc

import (
	"github.com/ferza17/ecommerce-microservices-v2/user-service/model/pb"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/presenter"
	userPostgresqlRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/repository/postgresql"
	userUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/user/usecase"
	eventStoreRepository "github.com/ferza17/ecommerce-microservices-v2/user-service/module/userEventStore/repository/mongodb"
	eventStoreUseCase "github.com/ferza17/ecommerce-microservices-v2/user-service/module/userEventStore/usecase"
)

func (srv *Server) RegisterService() {
	// Register Module

	newUserEventStoreRepository := eventStoreRepository.NewEventStoreRepository(srv.mongoDBConnector, srv.logger)
	newUserEventStoreUseCase := eventStoreUseCase.NewUserEventStoreUseCase(newUserEventStoreRepository, srv.logger)

	newUserPostgresqlRepository := userPostgresqlRepository.NewUserPostgresqlRepository(srv.postgresqlConnector, srv.logger)
	newUserUseCase := userUseCase.NewUserUseCase(newUserPostgresqlRepository, newUserEventStoreUseCase, srv.logger)

	pb.RegisterUserServiceServer(srv.grpcServer, presenter.NewUserPresenter(newUserUseCase))

}
