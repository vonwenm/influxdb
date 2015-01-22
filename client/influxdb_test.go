package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/influxdb/influxdb/client"
)

func TestNewClient(t *testing.T) {
	config := client.Config{}
	_, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
}

func TestNewClient_Defaults(t *testing.T) {
	config := client.Config{}
	c, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}

	defaultAddr := "localhost:8086"
	if c.Addr() != defaultAddr {
		t.Fatalf("unexpected addr: expected: %q, actual %q", defaultAddr, c.Addr())
	}
}

func TestClient_Ping(t *testing.T) {
	ts := emptyTestServer()
	defer ts.Close()

	config := client.Config{Addr: ts.URL}
	c, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
	d, err := c.Ping()
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
	if d == 0 {
		t.Fatalf("expected a duration greater than zero.  actual %v", d)
	}

}

func TestClient_Query(t *testing.T) {
	config := client.Config{}
	c, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}

	query := client.Query{}
	_, err = c.Query(query)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
}

func TestClient_Write(t *testing.T) {
	config := client.Config{}
	c, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}

	write := client.Write{}
	_, err = c.Write(write)
	if err != nil {
		t.Fatalf("unexpected error.  expected %v, actual %v", nil, err)
	}
}

// helper functions

func emptyTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		return
	}))
}