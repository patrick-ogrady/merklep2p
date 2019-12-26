package merklep2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateHashChildren(t *testing.T) {
	n := Node{
		Children: [][]byte{
			[]byte("left"),
			[]byte("right"),
		},
	}

	hash, err := n.CalculateHash()

	assert.NoError(t, err)
	assert.Equal(t, []byte("sick"), hash)
}