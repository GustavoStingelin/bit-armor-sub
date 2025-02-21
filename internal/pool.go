package internal

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"sync"
)

type ArmoredPool struct {
	outpoints map[chainhash.Hash]map[uint32]ArmoredOutpoint
	mu        *sync.RWMutex
}

var ap = &ArmoredPool{
	outpoints: make(map[chainhash.Hash]map[uint32]ArmoredOutpoint),
}

func PoolAdd(outpoint ArmoredOutpoint) {
	ap.mu.Lock()
	defer ap.mu.Unlock()

	if _, ok := ap.outpoints[outpoint.Hash]; !ok {
		ap.outpoints[outpoint.Hash] = make(map[uint32]ArmoredOutpoint)
	}

	ap.outpoints[outpoint.Hash][outpoint.Index] = outpoint
}

func PoolGet(hash chainhash.Hash, index uint32) (ArmoredOutpoint, bool) {
	ap.mu.RLock()
	defer ap.mu.RUnlock()

	if _, ok := ap.outpoints[hash]; !ok {
		return ArmoredOutpoint{}, false
	}

	outpoint, ok := ap.outpoints[hash][index]
	return outpoint, ok
}
