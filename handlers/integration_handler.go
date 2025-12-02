package handlers

import (
	"gohighload/metrics"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics.MetricsHandler().ServeHTTP(w, r)
}
