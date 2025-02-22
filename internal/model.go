package internal

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"log/slog"
)

type ArmoredOutpoint struct {
	Hash               chainhash.Hash
	Index              uint32
	value              int64
	DestinationAddress string
	SpendablePreSigned []PreSignedTx
}

type PreSignedTx struct {
	Hash  chainhash.Hash
	Index uint32
	fee   int64
	raw   []byte
}

func (ao *ArmoredOutpoint) FindNextPreSignedTx(minFee int64) (PreSignedTx, bool) {
	for _, preSignedTx := range ao.SpendablePreSigned {
		if preSignedTx.fee > minFee {
			return preSignedTx, true
		}
	}
	slog.Warn("No spendable pre-signed transaction found", "outpoint", ao)
	return PreSignedTx{}, false
}
