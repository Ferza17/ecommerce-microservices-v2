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
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().ConfigServicePayment.HttpHost, config.Get().ConfigServicePayment.MetricHttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().ConfigServicePayment.HttpHost, config.Get().ConfigServicePayment.MetricHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
