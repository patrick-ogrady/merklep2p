package merklep2p

import (
	"context"
	"errors"

	"github.com/btcsuite/btcutil/base58"
)

type MemStorage struct {
	store map[string][]byte
}

func NewMemStorage() Storage {
	return &MemStorage{
		store: make(map[string][]byte),
	}
}

func (m *MemStorage) Put(ctx context.Context, data []byte) ([]byte, error) {
	hash, err := CalculateHash(data)
	if err != nil {
		return nil, err
	}

	m.store[base58.Encode(hash)] = data
	return hash, nil
}

func (m *MemStorage) Get(ctx context.Context, hash []byte) ([]byte, error) {
	if val, ok := m.store[base58.Encode(hash)]; ok {
		return val, nil
	}

	return nil, errors.New("hash not found")
}
