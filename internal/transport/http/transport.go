package httptransport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alicobanserver/internal/service"
)

// MakeHTTPHandler ...
func MakeHTTPHandler(ctx context.Context, s service.Service) http.Handler {
	mux := http.NewServeMux()
	// POST /start
	mux.HandleFunc("/start", makeStartHandleFunc(ctx, s))

	return mux
}

// StartRequest ...
type StartRequest struct{}

func makeStartHandleFunc(ctx context.Context, s service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var sr StartRequest
		if err := json.NewDecoder(r.Body).Decode(&sr); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		s.Start(ctx)
		w.WriteHeader(http.StatusOK)
	}
}
