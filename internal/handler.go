package internal

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/btcsuite/btcd/wire"
	"log/slog"
)

func handleRawTx(topic string, txBytes []byte, seqBytes []byte) {
	slog.Debug("Received topic: %s\n", topic)
	slog.Debug("Raw Transaction: %s\n", hex.EncodeToString(txBytes))

	if len(seqBytes) == 4 {
		sequence := binary.LittleEndian.Uint32(seqBytes)
		slog.Debug("Sequence Number: %d\n", sequence)
	}

	tx := &wire.MsgTx{}
	err := tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		slog.Error("Failed to parse transaction: %v\n", err)
		return
	}
	txHash := tx.TxHash()
	slog.Debug("Parsed Transaction: %x %+v\n", txHash, tx)

TxInLoop:
	for _, txIn := range tx.TxIn {
		outpoint := txIn.PreviousOutPoint
		armoredOutPoint, ok := PoolGet(outpoint.Hash, outpoint.Index)
		if !ok {
			continue
		}
		slog.Info("Found armored outpoint: %+v\n", armoredOutPoint)

		for _, preSignedTx := range armoredOutPoint.SpendablePreSigned {
			buff := bytes.NewBuffer([]byte{})
			err := tx.Serialize(buff)
			if err != nil {
				slog.Error("Failed to serialize transaction: %v\n", err)
				continue
			}

			if bytes.Equal(preSignedTx.raw, buff.Bytes()) {
				slog.Info("Found matching pre-signed transaction: %+v\n", preSignedTx)
				break TxInLoop
			}
		}
		slog.Warn("Unknown transaction spending armored outpoint, armor activated: %+v\n", armoredOutPoint)
		actualFee := armoredOutPoint.value
		for _, txOut := range tx.TxOut {
			actualFee -= txOut.Value
		}
		slog.Warn("Actual fee: %d\n", actualFee)

		preSignedTx, ok := armoredOutPoint.FindNextPreSignedTx(actualFee)
		if !ok {
			slog.Error("No pre-signed transaction found for fee: %d\n", actualFee)
			continue
		}

		res, err := sendTransaction(preSignedTx.raw)
		if err != nil {
			return
		}
		slog.Warn("Sent pre-signed transaction: %s\n", res)
	}
}
