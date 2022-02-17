package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactions(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controllers.GetTransactions(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.NotContains(t, w.Body.String(), "error")
}
