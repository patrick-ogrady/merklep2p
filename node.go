package merklep2p

import (
	"crypto/sha256"

	"gopkg.in/mgo.v2/bson"
)

type Node struct {
	Left  []byte ",omitempty"
	Right []byte ",omitempty"

	Content []byte ",omitempty"
}

func (n *Node) CalculateHash() ([]byte, error) {
	data, err := bson.Marshal(n)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
