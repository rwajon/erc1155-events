package tests

import (
	"testing"

	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func TestStringToInt(t *testing.T) {
	assert.Equal(t, utils.StringToInt("1"), int64(1))
	assert.Equal(t, utils.StringToInt("0"), int64(0))
	assert.NotEqual(t, utils.StringToInt(""), float64(1))
}

func TestStringToFloat(t *testing.T) {
	assert.Equal(t, utils.StringToFloat("1"), float64(1))
	assert.Equal(t, utils.StringToFloat("0"), float64(0))
	assert.NotEqual(t, utils.StringToFloat(""), int64(1))
}

func TestHexToInt(t *testing.T) {
	assert.NotEqual(t, utils.HexToInt("0x0001"), 1)
	assert.Equal(t, utils.HexToInt("0x0001"), int64(1))
	assert.Equal(t, utils.HexToInt(""), int64(0))
}

func TestHexToFloat(t *testing.T) {
	assert.NotEqual(t, utils.HexToFloat("0x0001"), 1)
	assert.Equal(t, utils.HexToFloat("0x0001"), float64(1))
	assert.Equal(t, utils.HexToFloat(""), float64(0))
}
