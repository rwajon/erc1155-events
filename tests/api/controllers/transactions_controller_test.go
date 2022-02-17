package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
	"github.com/rwajon/erc1155-events/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests.DeleteTransactions()
	tests.CreateTransaction()

	controllers.GetTransactions(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.NotContains(t, w.Body.String(), "error")

}

func TestGetTransactionsError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests.DeleteTransactions()
	controllers.GetTransactions(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	assert.NotContains(t, w.Body.String(), "data")

}

func TestGetOneTransaction(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tests.DeleteTransactions()
	txHash := tests.CreateTransaction()

	c.Params = []gin.Param{{Key: "hash", Value: txHash}}
	controllers.GetOneTransaction(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.NotContains(t, w.Body.String(), "error")

}

func TestGetOneTransactionError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "hash", Value: "xxxxxx"}}

	controllers.GetOneTransaction(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "error")
	assert.NotContains(t, w.Body.String(), "data")
}
