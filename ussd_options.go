package ussdgwapp

func (app *ussdAppImpl) WithHost(host string) UssdApp {
	if host != "" {
		app.host = host
	}
	return app
}

func (app *ussdAppImpl) WithPort(port int) UssdApp {
	app.port = port
	return app
}

func (app *ussdAppImpl) WithPath(path string) UssdApp {
	if path != "" {
		app.path = path
	}
	return app
}

func (app *ussdAppImpl) WithMethod(method string) UssdApp {
	if method != "" {
		app.incomingMethod = method
	}
	return app
}

func (app *ussdAppImpl) WithMetricsPath(path string) UssdApp {
	if path == "" {
		app.metricsPath = path
	}
	return app
}

func (app *ussdAppImpl) WithHealthPath(path string) UssdApp {
	if path == "" {
		app.healthPath = path
	}
	return app
}
