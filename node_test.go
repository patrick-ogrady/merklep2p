package merklep2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportImport(t *testing.T) {
	node := &Node{
		Children: [][]byte{[]byte("left"), []byte("right")},
	}

	nodeData := node.Bytes()
	recoveredNode, err := RestoreNode(nodeData)

	assert.NoError(t, err)
	assert.Equal(t, node, recoveredNode)
}
