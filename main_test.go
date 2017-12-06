package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//HTTP request that will be pass to handler
	//method, route, request body
	req, err := http.NewRequest("GET", "", nil)
	//In case of error: stop testing
	if err != nil {
		t.Fatal(err)
	}

	//Create an http recorder to act as target of http request
	recorder := httptest.NewRecorder()

	//Create an HTTP handler from our handler function as defined in main.go to test
	hf := http.HandlerFunc(handler)

	//Serve the HTTP Request to our recorder to execute the handler to test
	hf.ServeHTTP(recorder, req)

	//Check status code is that which we expect
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Check the response body is what we expect.
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}

}