package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	router.ServeHTTP(w, req)

	res, _ := json.Marshal(map[string]string{"message": "pong"})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(res), w.Body.String())
}
