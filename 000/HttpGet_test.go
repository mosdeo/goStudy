package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	// assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

func BenchmarkPingRoute(b *testing.B) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	for i := 0; i < b.N; i++ {
		//fmt.Println(i, b.N)
		router.ServeHTTP(w, req)
		router.ServeHTTP(w, req)
		router.ServeHTTP(w, req)
		router.ServeHTTP(w, req)
	}

	// assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}
