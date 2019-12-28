package merklep2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLeaves(t *testing.T) {
	chunkSize := 1024
	arbData := []byte(RandomString(chunkSize*10 + 2))

	nodes := createLeaves(arbData, uint64(chunkSize))
	recoveredData := recoverData(nodes)

	assert.Equal(t, arbData, recoveredData)
}
