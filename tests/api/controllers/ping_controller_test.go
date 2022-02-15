package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	controllers.Ping(c)

	res, _ := json.Marshal(map[string]string{"message": "pong"})
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(res), w.Body.String())
}
