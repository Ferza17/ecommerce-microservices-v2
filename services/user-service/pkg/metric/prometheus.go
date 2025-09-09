package metric

import "github.com/prometheus/client_golang/prometheus"

var (
	GrpcRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_service_grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "status"},
	)

	GrpcRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "user_service_grpc_request_duration_seconds",
			Help:    "Duration of gRPC requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_service_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "user_service_http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	RabbitmqMessagesPublished = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_service_rabbitmq_messages_published_total",
			Help: "Total number of RabbitMQ messages published",
		},
		[]string{"queue", "status"},
	)

	RabbitmqMessagesConsumed = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_service_rabbitmq_messages_consumed_total",
			Help: "Total number of RabbitMQ messages consumed",
		},
		[]string{"queue", "status"},
	)
)
