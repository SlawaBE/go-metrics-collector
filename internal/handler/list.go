package handler

import (
	"net/http"
	"strings"

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
	w.Header().Add("Content-Type", "text/html")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	list, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	var builder strings.Builder
	builder.WriteString("<ul>")
	for _, v := range list {
		builder.WriteString("<li>")
		builder.WriteString(v)
		builder.WriteString("</li>")
	}
	builder.WriteString("</ul>")
	w.Write([]byte(builder.String()))
}
