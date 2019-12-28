package merklep2p

import (
	"github.com/btcsuite/btcutil/base58"
)

type MemStorage struct {
	store map[string]*Node
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		store: make(map[string]*Node),
	}
}

func (m *MemStorage) Put(node *Node) error {
	hash, err := node.CalculateHash()
	if err != nil {
		return err
	}

	m.store[base58.Encode(hash)] = node
	return nil
}

func (m *MemStorage) Get(nodeHash []byte) *Node {
	if val, ok := m.store[base58.Encode(nodeHash)]; ok {
		return val
	}

	return nil
}
