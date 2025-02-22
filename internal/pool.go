package internal

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"log/slog"
	"sync"
	"time"
)

type ArmoredPool struct {
	outpoints map[chainhash.Hash]map[uint32]ArmoredOutpoint
	mu        sync.RWMutex
	updatedAt time.Time
}

var ap = &ArmoredPool{
	outpoints: make(map[chainhash.Hash]map[uint32]ArmoredOutpoint),
	updatedAt: time.Time{},
}

func add(outpoint ArmoredOutpoint) {
	if _, ok := ap.outpoints[outpoint.Hash]; !ok {
		ap.outpoints[outpoint.Hash] = make(map[uint32]ArmoredOutpoint)
	}
	ap.outpoints[outpoint.Hash][outpoint.Index] = outpoint
}

func get(hash chainhash.Hash, index uint32) (ArmoredOutpoint, bool) {
	if _, ok := ap.outpoints[hash]; !ok {
		return ArmoredOutpoint{}, false
	}
	outpoint, ok := ap.outpoints[hash][index]
	return outpoint, ok
}

func PoolGet(hash chainhash.Hash, index uint32) (ArmoredOutpoint, bool) {
	now := time.Now()
	ap.mu.Lock()
	defer ap.mu.Unlock()

	if now.Sub(ap.updatedAt) > 3*time.Second {
		ap.updatedAt = now
		outpoints, err := GetArmoredOutpoints()
		if err != nil {
			slog.Error("Failed to get armored outpoints", "err", err)
			return ArmoredOutpoint{}, false
		}
		for _, outpoint := range outpoints {
			add(outpoint)
			//todo: remove old outpoints too
		}
	}

	return get(hash, index)
}
