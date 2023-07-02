package ussdgwapp

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var defaultHandlerFunction = func(*UssdRequest) *UssdResponse {
	return &UssdResponse{
		Message: "Welcome to USSD app using the USSD GW App package, this is the default handler.",
		Type:    RESPONSE_TYPE_EXISTING,
	}
}

func (app *ussdAppImpl) registerHealthHandler(mw Middleware) {
	http.Handle(app.healthPath, mw.WrapHandler(app.healthPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})))
}

func (app *ussdAppImpl) registerMetricsHandler(mw Middleware) {
	http.Handle("/metrics", mw.WrapHandler("/metrics", promhttp.HandlerFor(app.metricsRegistry, promhttp.HandlerOpts{
		Timeout:           10 * time.Second,
		EnableOpenMetrics: true,
	})))
}

func (app *ussdAppImpl) registerUssdHandlerFunction(mw Middleware) {
	http.Handle(app.path, mw.WrapHandler(app.path, http.HandlerFunc(app.handleUssdRequest)))
}

func (app *ussdAppImpl) handleUssdRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	ussdRequest, err := extractUrlQueryParams(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "error": "Invalid request body"`))
		return
	}

	// Default handler function if no custom handler function is set
	if !app.hasSetCustomHandlerFunction {
		ussdResponse := defaultHandlerFunction(ussdRequest)
		ussdResponseBytes, _ := json.Marshal(ussdResponse)
		w.WriteHeader(http.StatusOK)
		w.Write(ussdResponseBytes)
		return
	}

	ussdResponse := app.customHandlerFunction(ussdRequest)
	if !ussdResponse.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{ "error": "Invalid response from application using custom handler function"`))
		return
	}

	ussdResponseBytes, _ := json.Marshal(ussdResponse)
	w.WriteHeader(http.StatusOK)
	w.Write(ussdResponseBytes)
}
