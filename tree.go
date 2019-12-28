package merklep2p

import (
	"errors"
)

func numLeaves(treeData []byte, chunkSize uint64) uint64 {
	numLeaves := uint64(len(treeData)) / chunkSize
	if uint64(len(treeData))%chunkSize != 0 {
		numLeaves++
	}

	return numLeaves
}

func createLeaves(treeData []byte, chunkSize uint64) []*Node {
	leaves := make([]*Node, numLeaves(treeData, chunkSize))
	index := 0
	cursor := uint64(0)
	for uint64(len(treeData)) > cursor {
		limit := cursor + chunkSize
		if limit > uint64(len(treeData)) {
			limit = uint64(len(treeData))
		}

		chunk := treeData[cursor:limit]
		nodeContent := make([]byte, len(chunk))
		copy(nodeContent, chunk)
		leaves[index] = &Node{
			Content: nodeContent,
		}

		cursor += chunkSize
		index++
	}

	return leaves
}

func recoverData(nodes []*Node) []byte {
	originalBytes := make([]byte, 0)
	for _, node := range nodes {
		originalBytes = append(originalBytes, node.Content...)
	}

	return originalBytes
}

func NewTree(treeData []byte, chunkSize uint64) ([]*Node, *Node, error) {
	if len(treeData) == 0 {
		return nil, nil, errors.New("cannont construct tree with no content")
	}

	return nil, nil, nil
}
