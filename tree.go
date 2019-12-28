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

func buildLevel(nodes []*Node) ([]*Node, error) {
	numLevelNodes := len(nodes) / 2
	if len(nodes)%2 == 1 {
		numLevelNodes++ // needed if odd number of leaves
	}

	levelNodes := make([]*Node, numLevelNodes)
	for i := 0; i < numLevelNodes; i++ {
		left := i * 2
		right := i*2 + 1
		if right >= len(nodes) {
			right = left
		}

		leftHash, err := nodes[left].CalculateHash()
		if err != nil {
			return nil, err
		}

		rightHash, err := nodes[right].CalculateHash()
		if err != nil {
			return nil, err
		}

		levelNodes[i] = &Node{
			Left:  leftHash,
			Right: rightHash,
		}
	}

	return levelNodes, nil
}

func NewTree(treeData []byte, chunkSize uint64) ([]*Node, *Node, error) {
	if len(treeData) == 0 {
		return nil, nil, errors.New("cannont construct tree with no content")
	}

	leaves := createLeaves(treeData, chunkSize)

	allNodes := make([]*Node, len(leaves))
	copy(allNodes, leaves)

	currNodes := leaves
	level := 0
	for len(currNodes) > 1 {
		levelNodes, err := buildLevel(currNodes)
		if err != nil {
			return nil, nil, err
		}

		currNodes = levelNodes
		allNodes = append(allNodes, levelNodes...)
		level++
	}

	return allNodes, currNodes[0], nil
}
