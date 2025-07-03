package http

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func ServeHttpPrometheusMetricCollector() error {
	handler := http.NewServeMux()
	handler.Handle("/v1/notification/metrics", promhttp.Handler())
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().NotificationServiceHttpHost, config.Get().NotificationServiceMetricHttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().NotificationServiceHttpHost, config.Get().NotificationServiceMetricHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
