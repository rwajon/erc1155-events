package tests

import (
	"testing"

	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func TestInterfaceToJson(t *testing.T) {
	data := map[string]string{"name": "John Smith"}
	assert.Equal(t, string(utils.InterfaceToJson(data)), "{\"name\":\"John Smith\"}")
	assert.Equal(t, string(utils.InterfaceToJson(func() {})), "")
	assert.Nil(t, utils.InterfaceToJson(func() {}))
}
