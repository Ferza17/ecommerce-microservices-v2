package graphql

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	bootstrap2 "github.com/ferza17/ecommerce-microservices-v2/api-gateway/bootstrap"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/enum"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/model/graph/gen"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/transport/graphql/middleware"
	"github.com/ferza17/ecommerce-microservices-v2/api-gateway/transport/graphql/resolver"
	"github.com/go-chi/chi/v5"
	chim "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type (
	GraphQLTransport struct {
		httpServer *http.Server
		bootstrap  *bootstrap2.Bootstrap
	}
	Option func(server *GraphQLTransport)
)

func NewServer(address, port string, option ...Option) *GraphQLTransport {
	s := &GraphQLTransport{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", address, port),
			Handler: nil,
		},
	}
	for _, o := range option {
		o(s)
	}
	return s
}

func (srv *GraphQLTransport) Serve() {
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

	//transport.

	r := chi.NewRouter()
	r.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "CONNECT", "TRACE", "HEAD", "PATCH"},
			AllowedHeaders:   []string{enum.AcceptHeader.String(), enum.AuthorizationHeader.String(), enum.ContentTypeHeader.String(), enum.XCSRFTokenHeader.String(), enum.XRequestIDHeader.String()},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
		chim.RequestID,
		chim.RealIP,
		chim.Recoverer,
		chim.NoCache,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger(srv.bootstrap.Logger),
		chim.Heartbeat("/ping"),
	)

	// HTTP GraphQLTransport Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Handle("/query", server)
		r.Handle("/docs", playground.Handler("GraphQL playground", "/api/v1/query"))
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		srv.bootstrap.Logger.Error(fmt.Sprintf("%s %s", method, route))
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		srv.bootstrap.Logger.Error(fmt.Sprintf("error walking chi router: %v", err))
	}

	srv.httpServer.Handler = r
	err := srv.httpServer.ListenAndServe()
	if err != nil {
		srv.bootstrap.Logger.Error(fmt.Sprintf("error starting http transport: %v", err))
		return
	}
}

func (srv *GraphQLTransport) Close() error {
	if err := srv.httpServer.Close(); err != nil {
		srv.bootstrap.Logger.Error(fmt.Sprintf("error closing http transport: %v", err))
		return err
	}
	return nil
}
