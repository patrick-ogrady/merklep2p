package merklep2p

import (
	"context"
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

func TestCreateRecoverLevel(t *testing.T) {
	chunkSize := 1024
	children := 2
	arbData := []byte(RandomString(chunkSize*10 + 2))
	memStorage := NewMemStorage()
	ctx := context.Background()

	nodes := createLeaves(arbData, uint64(chunkSize))
	levelHashes, err := storeNodes(ctx, nodes, memStorage)
	assert.NoError(t, err)

	levelNodes, err := buildLevel(levelHashes, uint64(children))
	assert.NoError(t, err)

	recoveredLeaves, err := recoverLevel(ctx, levelNodes, memStorage)
	assert.NoError(t, err)
	assert.Equal(t, arbData, recoverData(recoveredLeaves))
}

func TestCreateRecoverTree(t *testing.T) {
	chunkSize := 1024
	children := 2
	arbData := []byte(RandomString(chunkSize*10 + 2))
	memStorage := NewMemStorage()
	ctx := context.Background()

	root, err := NewTree(ctx, arbData, uint64(chunkSize), uint64(children), memStorage)
	assert.NoError(t, err)
	assert.Equal(t, "B5CuNQ58N5tXT7sqBmbabiJa4HgR9aiwCcvCm9kpDRvc", base58.Encode(root))

	recoveredData, err := RecoverTree(ctx, root, memStorage)
	assert.NoError(t, err)
	assert.Equal(t, arbData, recoveredData)
}
