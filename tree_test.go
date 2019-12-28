package merklep2p

import (
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"
)

func TestCreateLeaves(t *testing.T) {
	chunkSize := 1024
	arbData := []byte(RandomString(chunkSize*10 + 2))

	nodes := createLeaves(arbData, uint64(chunkSize))
	recoveredData := recoverData(nodes)

	assert.Equal(t, arbData, recoveredData)
}

func TestBuildLevel(t *testing.T) {
	chunkSize := 1024
	arbData := []byte(RandomString(chunkSize*10 + 2))

	nodes := createLeaves(arbData, uint64(chunkSize))
	levelNodes, err := buildLevel(nodes)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(levelNodes))
}

func TestNewTree(t *testing.T) {
	chunkSize := 1024
	arbData := []byte(RandomString(chunkSize*10 + 2))

	nodes, rootNode, err := NewTree(arbData, uint64(chunkSize))
	assert.NoError(t, err)
	assert.Equal(t, 23, len(nodes))
	rootHash, err := rootNode.CalculateHash()
	assert.NoError(t, err)
	assert.Equal(t, "5nzGdJMc7vU17k7Mkgw2RTZHveKs2RVVzVTuMzfeq5i6", base58.Encode(rootHash))
}
