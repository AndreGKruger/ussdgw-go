package ussdgwapp

import (
	"testing"
)

func TestNewUssdApp(t *testing.T) {
	app := NewUssdApp("test")
	if app == nil {
		t.Error("NewUssdApp() should not return nil")
	}
}

func TestUssdAppImpl_WithHost(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).host != "0.0.0.0" {
		t.Error("WithHost() should set the default host")
	}

	// Test setting value
	app = NewUssdApp("test").WithHost("127.0.0.1")
	if app.(*ussdAppImpl).host != "127.0.0.1" {
		t.Error("WithHost() should set the host")
	}
}

func TestUssdAppImpl_WithPort(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).port != 80 {
		t.Error("WithPort() should set the default port")
	}

	// Test setting value
	app = NewUssdApp("test").WithPort(8080)
	if app.(*ussdAppImpl).port != 8080 {
		t.Error("WithPort() should set the port")
	}
}

func TestUssdAppImpl_WithPath(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).path != "/ussd" {
		t.Error("WithPath() should set the default path")
	}

	// Test setting value
	app = NewUssdApp("test").WithPath("/ussd2")
	if app.(*ussdAppImpl).path != "/ussd2" {
		t.Error("WithPath() should set the path")
	}
}

func TestUssdAppImpl_WithMethod(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).incomingMethod != "GET" {
		t.Error("WithMethod() should set the default method")
	}

	// Test setting value
	app = NewUssdApp("test").WithMethod("POST")
	if app.(*ussdAppImpl).incomingMethod != "POST" {
		t.Error("WithMethod() should set the method")
	}
}

func TestUssdAppImpl_WithMetricsPath(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).metricsPath != "/metrics" {
		t.Error("WithMetricsPath() should set the default metrics path")
	}

	// Test setting value
	app = NewUssdApp("test").WithMetricsPath("/metrics2")
	if app.(*ussdAppImpl).metricsPath != "/metrics2" {
		t.Error("WithMetricsPath() should set the metrics path")
	}
}

func TestUssdAppImpl_WithHealthPath(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).healthPath != "/health" {
		t.Error("WithHealthPath() should set the default health path")
	}

	// Test setting value
	app = NewUssdApp("test").WithHealthPath("/health2")
	if app.(*ussdAppImpl).healthPath != "/health2" {
		t.Error("WithHealthPath() should set the health path")
	}
}

func TestUssdAppImpl_WithAll(t *testing.T) {
	// Test default value
	app := NewUssdApp("test")
	if app.(*ussdAppImpl).host != "0.0.0.0" {
		t.Error("WithAll() should set the default host")
	}
	if app.(*ussdAppImpl).port != 80 {
		t.Error("WithAll() should set the default port")
	}
	if app.(*ussdAppImpl).path != "/ussd" {
		t.Error("WithAll() should set the default path")
	}
	if app.(*ussdAppImpl).incomingMethod != "GET" {
		t.Error("WithAll() should set the default method")
	}
	if app.(*ussdAppImpl).metricsPath != "/metrics" {
		t.Error("WithAll() should set the default metrics path")
	}
	if app.(*ussdAppImpl).healthPath != "/health" {
		t.Error("WithAll() should set the default health path")
	}

	// Test setting value
	app = NewUssdApp("test").WithHost("127.0.0.1").WithPort(8080).WithPath("/ussd2").WithMethod("POST").WithMetricsPath("/metrics2").WithHealthPath("/health2")
	if app.(*ussdAppImpl).host != "127.0.0.1" {
		t.Error("WithAll() should set the host")
	}
	if app.(*ussdAppImpl).port != 8080 {
		t.Error("WithAll() should set the port")
	}
	if app.(*ussdAppImpl).path != "/ussd2" {
		t.Error("WithAll() should set the path")
	}
	if app.(*ussdAppImpl).incomingMethod != "POST" {
		t.Error("WithAll() should set the method")
	}
	if app.(*ussdAppImpl).metricsPath != "/metrics2" {
		t.Error("WithAll() should set the metrics path")
	}
	if app.(*ussdAppImpl).healthPath != "/health2" {
		t.Error("WithAll() should set the health path")
	}

}
