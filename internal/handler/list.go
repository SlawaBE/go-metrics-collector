package handler

import (
	"net/http"

	"github.com/SlawaBE/go-metrics-collector/internal/service"
)

type ListMetricHandler struct {
	service *service.MetricsService
}

func NewListMetricHandler(service *service.MetricsService) *ListMetricHandler {
	return &ListMetricHandler{
		service: service,
	}
}

func (h *ListMetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	list := h.service.List()

	w.WriteHeader(http.StatusOK)
	for _, v := range list {
		w.Write([]byte(v + "\n"))
	}
}
