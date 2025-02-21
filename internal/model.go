package internal

import "github.com/btcsuite/btcd/wire"

type ArmoredOutpoint struct {
	wire.OutPoint
	value              int64
	DestinationAddress string
	SpendablePreSigned []PreSignedTx
}

type PreSignedTx struct {
	wire.OutPoint
	fee int64
	raw []byte
}

func (ao *ArmoredOutpoint) FindNextPreSignedTx(minFee int64) (PreSignedTx, bool) {
	for _, preSignedTx := range ao.SpendablePreSigned {
		if preSignedTx.fee > minFee {
			return preSignedTx, true
		}
	}
	return PreSignedTx{}, false
}
