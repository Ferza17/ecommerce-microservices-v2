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
		w.WriteHeader(http.StatusOK)
		return
	})
	log.Printf("Starting HTTP Metric Server on %s:%s", config.Get().NotificationServiceHttpHost, config.Get().NotificationServiceHttpPort)
	if err := http.ListenAndServe(
		fmt.Sprintf("%s:%s", config.Get().NotificationServiceHttpHost, config.Get().NotificationServiceHttpPort),
		handler,
	); err != nil {
		return err
	}
	return nil
}
