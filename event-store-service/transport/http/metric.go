package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ferza17/ecommerce-microservices-v2/event-store-service/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ServeHttpPrometheusMetricCollector() error {
	handler := http.NewServeMux()
	handler.Handle("/v1/event-store/metrics", promhttp.Handler())
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().EventStoreServiceHttpHost, config.Get().EventStoreServiceMetricHttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().EventStoreServiceHttpHost, config.Get().EventStoreServiceMetricHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
