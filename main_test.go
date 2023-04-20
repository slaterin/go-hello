package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETHome(t *testing.T) {
	t.Run("home returns 404", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			fmt.Println(err)
		}

		response := httptest.NewRecorder()

		home(response, request)

		got := response.Code
		want := http.StatusNotFound

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestGETHello(t *testing.T) {
	t.Run("returns hello", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/hello", nil)
		if err != nil {
			fmt.Println(err)
		}

		response := httptest.NewRecorder()

		hello(response, request)

		got := response.Body.String()
		want := "hello\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestGETHeaders(t *testing.T) {
	t.Run("headers gets a 200 status code", func(t *testing.T) {
		request, err := http.NewRequest(http.MethodGet, "/headers", nil)
		if err != nil {
			fmt.Println(err)
		}
		response := httptest.NewRecorder()

		request.Header.Add("Accept-Language", "json")
		headers(response, request)

		got := response.Body.String()
		want := "Accept-Language: json\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
