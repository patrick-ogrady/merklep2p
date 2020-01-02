package merklep2p

import (
	"bytes"
	"context"
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

func buildLevel(nodeHashes [][]byte) ([]*Node, error) {
	numLevelNodes := len(nodeHashes) / 2
	if len(nodeHashes)%2 == 1 {
		numLevelNodes++ // needed if odd number of leaves
	}

	levelNodes := make([]*Node, numLevelNodes)
	for i := 0; i < numLevelNodes; i++ {
		left := i * 2
		right := i*2 + 1
		if right >= len(nodeHashes) {
			right = left
		}

		levelNodes[i] = &Node{
			Left:  nodeHashes[left],
			Right: nodeHashes[right],
		}
	}

	return levelNodes, nil
}

func storeNodes(ctx context.Context, nodes []*Node, storage Storage) ([][]byte, error) {
	nodeHashes := make([][]byte, len(nodes))
	for i, node := range nodes {
		nodeHash, err := storage.Put(ctx, node.Bytes())
		if err != nil {
			return nil, err
		}

		nodeHashes[i] = nodeHash
	}

	return nodeHashes, nil
}

func NewTree(ctx context.Context, treeData []byte, chunkSize uint64, storage Storage) ([]byte, error) {
	if len(treeData) == 0 {
		return nil, errors.New("cannont construct tree with no content")
	}

	leaves := createLeaves(treeData, chunkSize)
	currHashes, err := storeNodes(ctx, leaves, storage)
	if err != nil {
		return nil, err
	}

	level := 0
	for len(currHashes) > 1 {
		levelNodes, err := buildLevel(currHashes)
		if err != nil {
			return nil, err
		}

		currHashes, err = storeNodes(ctx, levelNodes, storage)
		if err != nil {
			return nil, err
		}

		level++
	}

	return currHashes[0], nil
}

func nodeFromHash(ctx context.Context, hash []byte, storage Storage) (*Node, error) {
	nodeData, err := storage.Get(ctx, hash)
	if err != nil {
		return nil, err
	}

	node, err := RestoreNode(nodeData)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func levelHashes(nodes []*Node) [][]byte {
	levelHashes := make([][]byte, 0)
	for _, node := range nodes {
		levelHashes = append(levelHashes, node.Left)
		if bytes.Compare(node.Left, node.Right) == 0 {
			continue
		}

		levelHashes = append(levelHashes, node.Right)
	}

	return levelHashes
}

func recoverLevel(ctx context.Context, nodes []*Node, storage Storage) ([]*Node, error) {
	nextLevelHashes := levelHashes(nodes)
	nextLevelNodes := make([]*Node, len(nextLevelHashes))

	// TODO: Add multithreading to recovery
	for i, nodeHash := range nextLevelHashes {
		nodeData, err := storage.Get(ctx, nodeHash)
		if err != nil {
			return nil, err
		}

		node, err := RestoreNode(nodeData)
		if err != nil {
			return nil, err
		}

		nextLevelNodes[i] = node
	}

	return nextLevelNodes, nil
}

func RecoverTree(ctx context.Context, root []byte, storage Storage) ([]byte, error) {
	rootNode, err := nodeFromHash(ctx, root, storage)
	if err != nil {
		return nil, err
	}

	currNodes := []*Node{rootNode}
	for currNodes[0].Content == nil {
		currNodes, err = recoverLevel(ctx, currNodes, storage)
		if err != nil {
			return nil, err
		}
	}

	return recoverData(currNodes), nil
}
