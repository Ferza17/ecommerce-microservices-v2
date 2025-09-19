package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServeHttpPrometheusMetricCollector() error {
	handler := http.NewServeMux()
	handler.Handle("/v1/notification/metrics", promhttp.Handler())
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().ConfigServiceNotification.HttpHost, config.Get().ConfigServiceNotification.MetricHttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().ConfigServiceNotification.HttpHost, config.Get().ConfigServiceNotification.MetricHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
