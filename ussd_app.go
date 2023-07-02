package ussdgwapp

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type UssdApp interface {
	Start() error
	RegisterCustomHandlerFunction(handlerFunction func(*UssdRequest) *UssdResponse) UssdApp
	WithHost(host string) UssdApp
	WithPort(port int) UssdApp
	WithPath(path string) UssdApp
	WithMethod(method string) UssdApp
	WithMetricsPath(path string) UssdApp
	WithHealthPath(path string) UssdApp
	Close() error
}

type ussdAppImpl struct {
	name                        string
	host                        string
	port                        int
	path                        string
	customHandlerFunction       func(*UssdRequest) *UssdResponse
	hasSetCustomHandlerFunction bool
	incomingMethod              string
	metricsPath                 string
	metricsRegistry             *prometheus.Registry
	healthPath                  string
}

func NewUssdApp(name string) UssdApp {
	return &ussdAppImpl{
		name:                        name,
		host:                        "0.0.0.0",                // default host
		port:                        80,                       // default port
		path:                        "/ussd",                  // default path
		hasSetCustomHandlerFunction: false,                    // default hasSetCustomHandlerFunction
		incomingMethod:              "GET",                    // default incoming method
		metricsPath:                 "/metrics",               // default metrics path
		metricsRegistry:             prometheus.NewRegistry(), // default metrics registry
		healthPath:                  "/health",                // default health path
	}
}

func (app *ussdAppImpl) Start() error {
	if !app.hasSetCustomHandlerFunction {
		fmt.Println("WARNING: !!! No handler function has been set, using default handler function. Call RegisterHandlerFunction to register a custom handler. !!!")
	}

	host := fmt.Sprintf("%s:%v", app.host, app.port)

	// Register standard Go metrics and the process collector
	app.metricsRegistry.MustRegister(collectors.NewGoCollector(), collectors.NewProcessCollector(collectors.ProcessCollectorOpts{Namespace: transformToSnakeCase(app.name)}))
	// Register custom metrics middleware
	mw := newMiddleWare(app.metricsRegistry, nil)

	// Register handlers
	app.registerHealthHandler(mw)
	app.registerMetricsHandler(mw)
	app.registerUssdHandlerFunction(mw)

	printStartupMessage(app)
	return http.ListenAndServe(host, nil)
}

func (app *ussdAppImpl) RegisterCustomHandlerFunction(handlerFunction func(*UssdRequest) *UssdResponse) UssdApp {
	app.customHandlerFunction = handlerFunction
	app.hasSetCustomHandlerFunction = true
	return app
}

func (app *ussdAppImpl) Close() error {
	return nil
}
