package merklep2p

import (
	"bytes"
	"context"
	"io/ioutil"

	icore "github.com/ipfs/interface-go-ipfs-core"
	path "github.com/ipfs/interface-go-ipfs-core/path"
)

type IpfsStorage struct {
	api icore.CoreAPI
}

func NewIpfsStorage(api icore.CoreAPI) Storage {
	return &IpfsStorage{
		api: api,
	}
}

func (i *IpfsStorage) Put(ctx context.Context, data []byte) ([]byte, error) {
	blockStat, err := i.api.Block().Put(ctx, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return []byte(blockStat.Path().String()), nil
}

func (i *IpfsStorage) Get(ctx context.Context, hash []byte) ([]byte, error) {
	reader, err := i.api.Block().Get(ctx, path.New(string(hash)))
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return content, nil
}
