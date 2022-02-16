package tests

import (
	"testing"

	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func TestHexToInt(t *testing.T) {
	assert.NotEqual(t, utils.HexToInt("0x0001"), 1)
	assert.Equal(t, utils.HexToInt("0x0001"), int64(1))
	assert.Equal(t, utils.HexToInt(""), int64(0))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, utils.StringToInt("1"), int64(1))
	assert.Equal(t, utils.StringToInt("0"), int64(0))
	assert.NotEqual(t, utils.StringToInt(""), float64(1))
}
