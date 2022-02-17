package tests

import (
	"testing"

	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func TestJsonify(t *testing.T) {
	data := map[string]string{"name": "John Smith"}
	assert.Equal(t, string(utils.Jsonify(data)), "{\"name\":\"John Smith\"}")
	assert.Equal(t, string(utils.Jsonify(func() {})), "")
	assert.Nil(t, utils.Jsonify(func() {}))
}
