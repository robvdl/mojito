package mojito

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatic(t *testing.T) {
	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	m := New(&Options{})
	m.Use(NewStatic(http.Dir(".")))

	req, err := http.NewRequest("GET", "http://localhost:3000/mojito.go", nil)
	if err != nil {
		t.Error(err)
	}
	m.ServeHTTP(response, req)
	expect(t, response.Code, http.StatusOK)
	expect(t, response.Header().Get("Expires"), "")
	if response.Body.Len() == 0 {
		t.Errorf("Got empty body for GET request")
	}
}

func TestStaticHead(t *testing.T) {
	response := httptest.NewRecorder()
	response.Body = new(bytes.Buffer)

	m := New(&Options{})
	m.Use(NewStatic(http.Dir(".")))
	m.UseHandler(http.NotFoundHandler())

	req, err := http.NewRequest("HEAD", "http://localhost:3000/mojito.go", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(response, req)
	expect(t, response.Code, http.StatusOK)
	if response.Body.Len() != 0 {
		t.Errorf("Got non-empty body for HEAD request")
	}
}

func TestStaticAsPost(t *testing.T) {
	response := httptest.NewRecorder()

	m := New(&Options{})
	m.Use(NewStatic(http.Dir(".")))
	m.UseHandler(http.NotFoundHandler())

	req, err := http.NewRequest("POST", "http://localhost:3000/mojito.go", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(response, req)
	expect(t, response.Code, http.StatusNotFound)
}

func TestStaticBadDir(t *testing.T) {
	response := httptest.NewRecorder()

	m := Classic()
	m.UseHandler(http.NotFoundHandler())

	req, err := http.NewRequest("GET", "http://localhost:3000/mojito.go", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(response, req)
	refute(t, response.Code, http.StatusOK)
}

func TestStaticOptionsServeIndex(t *testing.T) {
	response := httptest.NewRecorder()

	m := New(&Options{})
	s := NewStatic(http.Dir("."))
	s.IndexFile = "mojito.go"
	m.Use(s)

	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(response, req)
	expect(t, response.Code, http.StatusOK)
}

func TestStaticOptionsPrefix(t *testing.T) {
	response := httptest.NewRecorder()

	m := New(&Options{})
	s := NewStatic(http.Dir("."))
	s.Prefix = "/public"
	m.Use(s)

	// Check file content behaviour
	req, err := http.NewRequest("GET", "http://localhost:3000/public/negroni.go", nil)
	if err != nil {
		t.Error(err)
	}

	m.ServeHTTP(response, req)
	expect(t, response.Code, http.StatusOK)
}
