package merklep2p

type Storage interface {
	Put([]byte) ([]byte, error)
	Get([]byte) ([]byte, error)
}
