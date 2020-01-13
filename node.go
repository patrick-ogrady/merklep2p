package merklep2p

import (
	"gopkg.in/mgo.v2/bson"
)

type Node struct {
	Children [][]byte ",omitempty"

	Content []byte ",omitempty"
}

func RestoreNode(data []byte) (*Node, error) {
	node := &Node{}
	err := bson.Unmarshal(data, node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (n *Node) Bytes() []byte {
	data, err := bson.Marshal(n)
	if err != nil {
		panic(err)
	}

	return data
}
