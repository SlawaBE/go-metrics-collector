package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/SlawaBE/go-metrics-collector/internal/retry"
)

type PingHandler struct {
	db *sql.DB
}

func NewPingHandler(db *sql.DB) *PingHandler {
	return &PingHandler{
		db: db,
	}
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		w.WriteHeader(http.StatusOK)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	err := retry.RetryWithBackoff(ctx, func() error {
		return h.db.PingContext(ctx)
	})
	if err != nil {
		http.Error(w, "No connect to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
