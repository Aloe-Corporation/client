package test

import (
	"net/http"
	"net/http/httptest"
)

var (
	PingEndpointServerURL = "http://server.com"
)

func GetPingEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" || r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Status not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func GetEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/get" || r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is data"))
	}))
}

func GetEndpointWithoutBody() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/get" || r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func GetEndpointWithHeader() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/get" || r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}

		if r.Header.Get("test-header") == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Missing test-header Header in request."))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is data"))
	}))
}

func PostEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/post" || r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is data"))
	}))
}

func PutEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/put" || r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is data"))
	}))
}

func DeleteEndpoint() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/delete" || r.Method != "DELETE" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is data"))
	}))
}
