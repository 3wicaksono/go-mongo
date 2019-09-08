package helpers

import (
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Run("Test helper api response", func(t *testing.T) {
		rw := httptest.NewRecorder()
		Response(rw, 200, nil)
	})

	t.Run("Test helper api response", func(t *testing.T) {
		rw := httptest.NewRecorder()
		Response(rw, 0, nil)
	})
}
