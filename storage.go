package merklep2p

type Storage interface {
	Put(*Node) error
	Get([]byte) *Node
}
