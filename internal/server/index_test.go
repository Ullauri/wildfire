package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestIndexHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	root := router.Group("")
	{
		AddIndexRoutes(root, ServerTest.NamesClient, ServerTest.JokesClient)
	}

	t.Run("200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("expected status OK; got %v", w.Code)
		}

		expected := "This is a joke"
		if w.Body.String() != expected {
			t.Errorf("expected body %q; got %q", expected, w.Body.String())
		}
	})

	t.Run("500", func(t *testing.T) {
		// reset test state
		defer func() {
			ServerTest.NamesClient.GetRandomNameError = nil
		}()
		ServerTest.NamesClient.GetRandomNameError = fmt.Errorf("some error")

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("expected status OK; got %v", w.Code)
		}
	})
}
