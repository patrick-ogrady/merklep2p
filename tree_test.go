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
	memStorage := NewMemStorage()

	nodes := createLeaves(arbData, uint64(chunkSize))
	levelHashes, err := storeNodes(nodes, memStorage)
	assert.NoError(t, err)

	levelNodes, err := buildLevel(levelHashes)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(levelNodes))
}

func TestCreateRecoverTree(t *testing.T) {
	chunkSize := 1024
	arbData := []byte(RandomString(chunkSize*10 + 2))
	memStorage := NewMemStorage()

	root, err := NewTree(arbData, uint64(chunkSize), memStorage)
	assert.NoError(t, err)
	assert.Equal(t, "5nzGdJMc7vU17k7Mkgw2RTZHveKs2RVVzVTuMzfeq5i6", base58.Encode(root))

	recoveredData, err := RecoverTree(root, memStorage)
	assert.NoError(t, err)
	assert.Equal(t, arbData, recoveredData)
}
