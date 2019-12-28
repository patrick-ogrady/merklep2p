package merklep2p

import (
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"
)

func TestCalculateHash(t *testing.T) {
	n := Node{
		Left:  []byte("left"),
		Right: []byte("right"),
	}

	hash, err := n.CalculateHash()

	assert.NoError(t, err)
	assert.Equal(t, "3VNzeWnMPADdMM2fcfrQxWex2nHupwvvGA1egGwaSKFM", base58.Encode(hash))
}
