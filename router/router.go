package router

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LogPayload(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost || r.Method == http.MethodPut {

			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "cannot read request body", http.StatusBadRequest)
				return
			}

			log.Printf("POST %s payload: %s", r.URL.Path, string(bodyBytes))

			// restore body so next handler can read it
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		next.ServeHTTP(w, r)
	})
}

func MainRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(LogPayload)
	r.Mount("/power", powerRouter())
	return r
}
