package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	s := Hello()
	if s != "Hello, world." {
		t.Errorf("Hello() returns incorrect result")
	}
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Hello, world.\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello()
	}
}

func ExampleHello() {
	fmt.Println(Hello())
	// Output: Hello, world.
}
