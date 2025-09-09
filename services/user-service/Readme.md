**USER SERVICE**
---

**Folder Structure**


```bash

./src
├── cmd                                 -> App Command Line entrypoints (e.g., run, migrate, etc.)
├── config                              -> App configuration loaders (e.g., from Consul, env files)
├── infrastructure                      -> Initialize App Infrastructure [Database, Message Broker, External Services via gRPC/HTTP, Telemetry (Jaeger, OpenTelemetry)]
├── interceptor                         -> App Interceptors [gRPC & HTTP Middleware such as Logging, Tracing, Auth]
├── main.go                             -> Main entry point of the application
├── model                               -> App Domain Models
│   ├── orm                             -> ORM Library integration with Gorm
│   └── rpc                             -> Protobuf-generated gRPC server & client code
├── module                              -> Application Modules (Clean Architecture structure)
│   ├── <feature_name>                  -> Feature module (e.g., access_control, auth, role, user)
│   │   ├── http_presenter              -> HTTP handler/controller layer
│   │   ├── grpc_presenter              -> gRPC service implementations
│   │   ├── usecase                     -> Business logic layer
│   │   └── repository                  -> Data layer (e.g., PostgreSQL, Redis)
├── package                             -> Shared application packages/utilities
│   ├── context                         -> Request context handling (e.g., auth, request ID)
│   └── worker_pool                     -> Worker pool implementation for background tasks
│   └── metrics                         -> Metrics implementation for metric collector using prometheus
├── transport                           -> Application-level transport setup [gRPC Server, HTTP Server, Message Broker like RabbitMQ/Kafka]
└── util                                -> App-wide utilities and helpers (e.g., string, time, logging helpers)
```
---

**TRANSPORT**
- RABBITMQ
- HTTP
- GRPC
---

**METRIC**
- PROMETHEUS
---

**MONITORING**
- JAEGER
---

**Config**
- ENV       : [env.local, env.production, env.example]
- CONSUL
---