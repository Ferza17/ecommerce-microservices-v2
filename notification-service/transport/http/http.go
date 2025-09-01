package http

import (
	"fmt"
	"github.com/ferza17/ecommerce-microservices-v2/notification-service/config"
	"log"
	"net/http"
)

func ServeHttpHealthCheckHandler() error {
	handler := http.NewServeMux()
	handler.HandleFunc("/v1/notification/check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	addr := fmt.Sprintf("%s:%s", config.Get().NotificationServiceHttpHost, config.Get().NotificationServiceHttpPort)
	log.Printf("Starting HTTP HealthCheck Server on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil && err != http.ErrServerClosed {
		log.Fatalf("healthcheck server failed: %v", err)
		return err
	}
	return nil
}
