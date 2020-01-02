package merklep2p

import "context"

type Storage interface {
	Put(context.Context, []byte) ([]byte, error)
	Get(context.Context, []byte) ([]byte, error)
}
