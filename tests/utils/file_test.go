package tests

import (
	"testing"

	"github.com/rwajon/erc1155-events/utils"
	"github.com/stretchr/testify/assert"
)

func TestWriteToFile(t *testing.T) {
	assert.Nil(t, utils.WriteToFile("test.log", "testing..."))
	assert.NotNil(t, utils.WriteToFile("", "testing..."))
	assert.NotNil(t, utils.WriteToFile("", ""))
}
