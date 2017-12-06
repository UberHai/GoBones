package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	//Instantiate the router using the constructor func in main
	r := newRouter()

	//Create a new server using the "httptest" libs `NewServer` method
	mockServer := httptest.NewServer(r)

	//Mock server runs and exposes it's location in the URL attribute
	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	//Ensure status: 200(ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200: OK, got %d", resp.StatusCode)
	}

	//Read response body and convert to string
	defer resp.Body.Close()
	//Read as bytes
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	//Convert bytes to string
	respString := string(b)
	expected := "Hello World!"

	//Check if response matches the expected output
	if respString != expected {
		t.Errorf("Response should be \"%s\" got \"%s\"", expected, respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	//Create a request to a route we did NOT define "POST /hello"
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	//We want status to be 405(not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405 (Not Allowed), got \"%d\"", resp.StatusCode)
	}

	//Expect an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be \"%s\" got \"%s\"", expected, respString)
	}
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	//We want to hit the `GET /assets/` router to get index.html
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	//We want our status to be 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: Status 200, Got \"%d\"", resp.StatusCode)
	}

	//Test that the content-type header is "text/html"; charset=utf-8
	//So we know an html file was server
	contentType := resp.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}
