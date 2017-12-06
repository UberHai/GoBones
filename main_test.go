package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
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
	if resp.StatusCode != http.StatusOK{
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