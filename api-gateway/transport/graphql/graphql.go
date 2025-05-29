package graphql

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	bootstrap2 "github.com/ferza17/ecommerce-microservices-v2/api-gateway/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/graph/gen"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/transport/graphql/middleware"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/transport/graphql/resolver"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type GraphQLTransport struct {
	httpServer *http.Server
	host, port string
	bootstrap  *bootstrap2.Dependency
}

func NewGraphQLTransport(host, port string, b *bootstrap2.Dependency) *GraphQLTransport {
	return &GraphQLTransport{
		host:      host,
		port:      port,
		bootstrap: b,
	}
}

func (srv *GraphQLTransport) Serve() {
	router := mux.NewRouter()
	server := handler.New(gen.NewExecutableSchema(gen.Config{
		Resolvers: &resolver.Resolver{
			UserUseCase:             srv.bootstrap.UserUseCase,
			ProductUseCase:          srv.bootstrap.ProductUseCase,
			CartUseCase:             srv.bootstrap.CartUseCase,
			AuthUseCase:             srv.bootstrap.AuthUseCase,
			TelemetryInfrastructure: srv.bootstrap.TelemetryInfrastructure,
		},
		Directives: gen.DirectiveRoot{
			//Jwt: middleware.DirectiveJwtRequired,
		},
	},
	))
	server.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	server.AddTransport(transport.Options{})
	server.AddTransport(transport.GET{})
	server.AddTransport(transport.POST{})
	server.AddTransport(transport.MultipartForm{})
	server.Use(extension.Introspection{})

	router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]any{"message": "pong"})
	}))
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.Use(middleware.Logger(srv.bootstrap.Logger))

	// Public
	public := apiV1.PathPrefix("/public").Subrouter()
	public.HandleFunc("/user/register", srv.bootstrap.AuthPresenter.CreateUser).Methods(http.MethodPost)
	public.HandleFunc("/user/login", srv.bootstrap.AuthPresenter.UserLoginByEmailAndPassword).Methods(http.MethodPost)
	public.HandleFunc("/user/logout", srv.bootstrap.AuthPresenter.UserLogoutByToken).Methods(http.MethodPost)
	public.HandleFunc("/user/verify_otp", srv.bootstrap.AuthPresenter.UserVerifyOtp).Methods(http.MethodPost)

	private := apiV1.PathPrefix("/private").Subrouter()
	private.Use(middleware.Authorization(srv.bootstrap.AuthServiceInfrastructure, srv.bootstrap.TelemetryInfrastructure))
	private.Handle("/query", server)
	private.Handle("/docs", playground.Handler("GraphQL playground", "/api/v1/private/query"))

	srv.httpServer = &http.Server{
		Addr:         srv.host + ":" + srv.port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	if err := srv.httpServer.ListenAndServe(); err != nil {
		srv.bootstrap.Logger.Error(fmt.Sprintf("err : %v", err))
	}
}

func (srv *GraphQLTransport) Close() error {
	if err := srv.httpServer.Close(); err != nil {
		srv.bootstrap.Logger.Error(fmt.Sprintf("error closing http transport: %v", err))
		return err
	}
	return nil
}
