# Ussd Application using the Ussd Gateway

## Basic features
 - Starts a webserver that will receive requests from the ussd gateway and respond to them
 - Register a customer function as a handler that will be called when the customer sends a request to the ussd gateway
 - Webserver can be configured 
 - Webserver has a default health and metrics endpoint
 - Promethues metrics will be exposed on the metrics endpoint for the requests received by the ussd gateway

## Usage

```go
    var customhandlerfunction = func(req *ussdgwapp.UssdRequest) *ussdgwapp.UssdResponse {
        return &ussdgwapp.UssdResponse{
            Type:    ussdgwapp.REQUEST_TYPE_EXISTING,
            Message: "Welcome to USSD app using the USSD GW App package, this is the custom handler.",
        }
    }

    func main(){
        ussdapp := ussdgwapp.NewUssdApp().WithHost("127.0.0.1").WithPort(8080).WithPath("/v1/ussd")
        ussdapp.RegisterCustomHandlerFunction(customhandlerfunction)
        err := ussdapp.Start()
        if err != nil {
            log.Fatal(err)
        }
    } 

```