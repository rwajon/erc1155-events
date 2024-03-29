package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rwajon/erc1155-events/api/routes"
	"github.com/rwajon/erc1155-events/tests"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routes.Init(tests.InitApp())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	res, _ := json.Marshal(map[string]string{"message": "pong"})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(res), w.Body.String())
}
