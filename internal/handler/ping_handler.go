package handler

import (
	"context"
	"database/sql"
	"net/http"
	"time"
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
	}
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	if err := h.db.PingContext(ctx); err != nil {
		http.Error(w, "No connect to database", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
