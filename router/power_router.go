package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/isubhampadhi56/remote-management/pkg/power"
)

type Request struct {
	Host string `json:"host"`
}

type QueryResponse struct {
	PowerState string `json:"power_state"`
}

func getHostFromRequest(r *http.Request) (string, error) {
	// if r.Method == http.MethodPost {
	// 	var req Request
	// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 		return "", fmt.Errorf("invalid request body: %v", err)
	// 	}
	// 	if req.Host == "" {
	// 		return "", fmt.Errorf("host required in request body")
	// 	}
	// 	return req.Host, nil
	// }

	// For GET requests, get host from URL parameter
	host := chi.URLParam(r, "host")
	if host == "" {
		return "", fmt.Errorf("host parameter required in URL")
	}
	return host, nil
}

func manager(host string) *power.Manager {
	return &power.Manager{
		Host:     host,
		Port:     "623",
		Username: "Administrator",
		Password: "Realtek",
	}
}

func powerOn(w http.ResponseWriter, r *http.Request) {
	host, err := getHostFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pm := manager(host)
	pm.PowerOn()

	w.WriteHeader(http.StatusOK)
}

func powerOff(w http.ResponseWriter, r *http.Request) {
	host, err := getHostFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pm := manager(host)
	pm.PowerOff()

	w.WriteHeader(http.StatusOK)
}

func powerCycle(w http.ResponseWriter, r *http.Request) {
	host, err := getHostFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pm := manager(host)
	pm.PowerCycle()

	w.WriteHeader(http.StatusOK)
}

func powerQuery(w http.ResponseWriter, r *http.Request) {
	host, err := getHostFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pm := manager(host)

	state, _ := pm.PowerState()

	resp := QueryResponse{
		PowerState: state.String(),
	}

	json.NewEncoder(w).Encode(resp)
}

func powerRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/on/{host}", powerOn)
	r.Get("/on/{host}", powerOn)

	r.Post("/off/{host}", powerOff)
	r.Get("/off/{host}", powerOff)

	r.Post("/cycle/{host}", powerCycle)
	r.Get("/cycle/{host}", powerCycle)

	r.Post("/status/{host}", powerQuery)
	r.Get("/status/{host}", powerQuery)

	return r
}
