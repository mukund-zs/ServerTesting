package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFun(t *testing.T) {
	const (
		url  = "localhost:8000/hello?name="
		name = "mukund"
	)
	testcases := []struct {
		method string
		name   string
		output int
	}{
		{http.MethodGet, name, http.StatusOK},
		{http.MethodPost, name, http.StatusMethodNotAllowed},
		{http.MethodPut, name, http.StatusMethodNotAllowed},
		{http.MethodDelete, name, http.StatusMethodNotAllowed},
		{http.MethodGet, "", http.StatusBadRequest},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(v.method, url+v.name, nil)
		w := httptest.NewRecorder()
		helloHandler(w, req)
		if w.Result().StatusCode != v.output {
			t.Errorf("Test Failed %v. Expected %v GOT %v", i, v.output, w.Result().StatusCode)
		}
	}
}
