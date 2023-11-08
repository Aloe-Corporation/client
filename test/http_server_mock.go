package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// GetPingEndpoint is a HTTP mock endpoint used for testing.
// The ping url is "/".
func GetPingEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Status not found")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
}

// GetEndpoint is a HTTP mock endpoint to simulate a GET request that responds with data.
// The only valid path is "/get".
func GetEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/get" || r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Status not found")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("This is data")); err != nil {
			fmt.Println("can't write in response writer: ", err.Error())
		}
	}))
}

// GetEndpointWithHeader is a HTTP mock endpoint to simulate a GET request that await for
// "test-header" header in the request. The server will respond with a 404 status code
// if the incoming request do not contains the header.
// The only valid path is "/get".
func GetEndpointWithHeader() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/get" || r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Status not found")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
			return
		}

		if r.Header.Get("test-header") == "" {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Missing test-header Header in request.")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("This is data")); err != nil {
			fmt.Println("can't write in response writer: ", err.Error())
		}
	}))
}

// PostEndpoint is a HTTP mock endpoint to simulate a POST request that responds without data.
// The only valid path is "/post".
func PostEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/post" || r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Status not found")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("This is data")); err != nil {
			fmt.Println("can't write in response writer: ", err.Error())
		}
	}))
}

// PutEndpoint is a HTTP mock endpoint to simulate a PUT request that responds without data.
// The only valid path is "/put".
func PutEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/put" || r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Status not found")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("This is data")); err != nil {
			fmt.Println("can't write in response writer: ", err.Error())
		}
	}))
}

// DeleteEndpoint is a HTTP mock endpoint to simulate a DELETE request that responds without data.
// The only valid path is "/delete".
func DeleteEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/delete" || r.Method != "DELETE" {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Status not found")); err != nil {
				fmt.Println("can't write in response writer: ", err.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("This is data")); err != nil {
			fmt.Println("can't write in response writer: ", err.Error())
		}
	}))
}
