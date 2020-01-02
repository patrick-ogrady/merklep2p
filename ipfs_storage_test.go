package merklep2p

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ipfs/go-datastore"
	syncds "github.com/ipfs/go-datastore/sync"
	"github.com/ipfs/go-filestore"
	"github.com/ipfs/go-ipfs-config"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	mock "github.com/ipfs/go-ipfs/core/mock"
	"github.com/ipfs/go-ipfs/keystore"
	"github.com/ipfs/go-ipfs/repo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p/p2p/net/mock"
	"github.com/stretchr/testify/assert"
)

func getMockAPI(ctx context.Context) (icore.CoreAPI, error) {
	mn := mocknet.New(ctx)
	c := config.Config{}
	c.Addresses.Swarm = []string{"/ip4/127.0.0.1/tcp/4001"}
	c.Identity = config.Identity{
		PeerID: "QmTFauExutTsy4XP6JbMFcw2Wa9645HJt2bTqL6qYDCKfe",
	}
	c.Experimental.FilestoreEnabled = true

	ds := syncds.MutexWrap(datastore.NewMapDatastore())
	r := &repo.Mock{
		C: c,
		D: ds,
		K: keystore.NewMemKeystore(),
		F: filestore.NewFileManager(ds, filepath.Dir(os.TempDir())),
	}

	node, err := core.NewNode(ctx, &core.BuildCfg{
		Repo:   r,
		Host:   mock.MockHostOption(mn),
		Online: false,
		ExtraOpts: map[string]bool{
			"pubsub": true,
		},
	})
	if err != nil {
		return nil, err
	}

	api, err := coreapi.NewCoreAPI(node)
	if err != nil {
		return nil, err
	}

	return api, nil
}

func TestPutGet(t *testing.T) {
	ctx := context.Background()
	api, err := getMockAPI(ctx)
	assert.NoError(t, err)
	arbData := []byte("hello")

	storage := NewIpfsStorage(api)
	path, err := storage.Put(ctx, arbData)
	assert.NoError(t, err)

	recoveredData, err := storage.Get(ctx, path)
	assert.NoError(t, err)

	assert.Equal(t, arbData, recoveredData)
}
