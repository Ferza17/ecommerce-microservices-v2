package http

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/user-service/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func ServeHTTPMetricCollector() {
	// Start metrics server
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Metrics server listening on %s:%s", config.Get().UserServiceHttpHost, config.Get().UserServiceHttpMetricPort)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Get().UserServiceHttpHost, config.Get().UserServiceHttpMetricPort), nil); err != nil {
		log.Fatalf("Failed to serve metrics: %v", err)
	}
}
