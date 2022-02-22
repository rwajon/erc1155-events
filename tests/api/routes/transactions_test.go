package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rwajon/erc1155-events/api/routes"
	"github.com/rwajon/erc1155-events/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionsRoute(t *testing.T) {
	tests.DeleteTransactions()
	tests.CreateTransaction()

	router := routes.Init(tests.InitApp())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/transactions/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.NotContains(t, w.Body.String(), "error")
}

func TestGetOneTransactionRoute(t *testing.T) {
	tests.DeleteTransactions()
	txHash := tests.CreateTransaction()

	router := routes.Init(tests.InitApp())

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/transactions/"+txHash, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.NotContains(t, w.Body.String(), "error")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/v1/transactions/xxxxxxxx", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	assert.NotContains(t, w.Body.String(), "data")
}
