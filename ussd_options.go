package ussdgwapp

func (app *ussdAppImpl) WithHost(host string) UssdApp {
	app.host = host
	return app
}

func (app *ussdAppImpl) WithPort(port int) UssdApp {
	app.port = port
	return app
}

func (app *ussdAppImpl) WithPath(path string) UssdApp {
	app.path = path
	return app
}

func (app *ussdAppImpl) WithMethod(method string) UssdApp {
	app.incomingMethod = method
	return app
}

func (app *ussdAppImpl) WithMetricsPath(path string) UssdApp {
	app.metricsPath = path
	return app
}

func (app *ussdAppImpl) WithHealthPath(path string) UssdApp {
	app.healthPath = path
	return app
}
