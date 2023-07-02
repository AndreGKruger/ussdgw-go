package ussdgwapp

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func transformToSnakeCase(input string) string {
	var output string
	for i, r := range input {
		if i > 0 && r >= 'A' && r <= 'Z' {
			output += "_"
		}
		output += string(r)
	}
	output = strings.Replace(output, "-", "_", -1)
	output = strings.ToLower(output)
	return output
}

func extractUrlQueryParams(r *http.Request) (*UssdRequest, error) {
	//msisdn
	msisdn := r.URL.Query().Get("msisdn")
	m, err := strconv.ParseUint(msisdn, 10, 64)
	if err != nil {
		return nil, errors.New(`{ "error": "Invalid msisdn, cannot parse to Uint64" }`)
	}
	//type
	r_type := r.URL.Query().Get("type")
	t, err := strconv.Atoi(r_type)
	if err != nil {
		return nil, errors.New(`{ "error": "Invalid type, cannot parse to int" }`)
	}
	//msg
	msg, err := url.QueryUnescape(r.URL.Query().Get("msg"))
	if err != nil {
		return nil, errors.New(`{ "error": "Invalid msg, cannot unescape" }`)
	}

	return &UssdRequest{
		SessionId: r.URL.Query().Get("sessionid"),
		Msisdn:    m,
		Type:      t,
		Message:   msg,
	}, nil
}

func printStartupMessage(app *ussdAppImpl) {
	fmt.Printf("%s - [INFO] - Started server for %s on - http://%s:%v \n", time.Now().Format(time.RFC3339), app.name, app.host, app.port)
	fmt.Printf("%s - [INFO] - Registered Health endpoint on - http://%s:%v%s \n", time.Now().Format(time.RFC3339), app.host, app.port, app.healthPath)
	fmt.Printf("%s - [INFO] - Registered Metrics endpoint on - http://%s:%v%s \n", time.Now().Format(time.RFC3339), app.host, app.port, app.metricsPath)
	fmt.Printf("%s - [INFO] - Registered Ussd endpoint on - http://%s:%v%s \n", time.Now().Format(time.RFC3339), app.host, app.port, app.path)
}
