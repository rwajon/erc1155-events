package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rwajon/erc1155-events/api/routes"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRoutes(t *testing.T) {
	router := routes.Init()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/transactions/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
