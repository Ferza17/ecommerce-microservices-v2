package http

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/payment-service/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func ServeHttpPrometheusMetricCollector() error {
	handler := http.NewServeMux()
	handler.Handle("/v1/payment/metrics", promhttp.Handler())
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().PaymentServiceHttpHost, config.Get().PaymentServiceMetricHttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().PaymentServiceHttpHost, config.Get().PaymentServiceMetricHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
