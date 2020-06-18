package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_handler(t *testing.T) {
	host, _ := os.Hostname()
	type args struct {
		path string
		status int
		expected string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"status OK",
			args{"/", http.StatusOK, fmt.Sprintf("Hello from %s", host)},
		},
	}

	rr := httptest.NewRecorder()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.args.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			handler := http.HandlerFunc(handler)
			handler.ServeHTTP(rr, req)
			if status := rr.Code; status != tt.args.status {
				t.Errorf("handler returned wrong status code: got [%v] want [%v]",
					status, tt.args.status)
				return
			}
			if rr.Body.String() != tt.args.expected {
				t.Errorf("handler returned unexpected body: got [%v] want [%v]",
					rr.Body.String(), tt.args.expected)
			}
		})
	}
}
