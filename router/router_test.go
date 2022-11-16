package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter_Route(t *testing.T) {
	method := http.MethodGet
	path := "/hello/world"
	h1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	router := NewRouter()
	router.Route(method, path, h1)
	route := router.root.traverse(path, false)

	h2 := route.handlerByMethod[method]
	if h2 == nil {
		t.Fatalf("Handler is not found for \"%s %s\"\n", method, path)
	}

	res := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, nil)
	h1.ServeHTTP(res, req)
	c1 := res.Code
	s1 := res.Body.String()

	res = httptest.NewRecorder()
	req = httptest.NewRequest(method, path, nil)
	h2.ServeHTTP(res, req)
	c2 := res.Code
	s2 := res.Body.String()

	if c1 != c2 {
		t.Fatalf("Got status code: %d, want: %d\n", c1, c2)
	}
	if s1 != s2 {
		t.Errorf("Got body: %s, want: %s\n", s1, s2)
	}
}
