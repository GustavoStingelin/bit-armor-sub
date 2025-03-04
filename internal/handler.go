package internal

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"log/slog"
)

func handleRawTx(topic string, txBytes []byte, seqBytes []byte) {
	slog.Debug("Received topic", "topic", topic)
	slog.Debug("Raw Transaction", "rawTx", hex.EncodeToString(txBytes))

	if len(seqBytes) == 4 {
		sequence := binary.LittleEndian.Uint32(seqBytes)
		slog.Debug("Sequence Number", "seq", sequence)
	}

	tx := &wire.MsgTx{}
	err := tx.Deserialize(bytes.NewReader(txBytes))
	if err != nil {
		slog.Error("Failed to parse transaction", "err", err)
		return
	}
	txId := tx.TxID()
	slog.Info("Parsed Transaction",
		"txId", txId,
		"tx", fmt.Sprintf("%+v", tx),
	)

TxInLoop:
	for _, txIn := range tx.TxIn {
		outpoint := txIn.PreviousOutPoint
		armoredOutPoint, ok := PoolGet(outpoint.Hash, outpoint.Index)
		if !ok {
			continue
		}
		slog.Info("Found armored outpoint", "outpoint", armoredOutPoint)

		for _, preSignedTx := range armoredOutPoint.SpendablePreSigned {
			buff := bytes.NewBuffer([]byte{})
			err := tx.Serialize(buff)
			if err != nil {
				slog.Error("Failed to serialize transaction", "err", err)
				continue
			}

			if bytes.Equal(preSignedTx.raw, buff.Bytes()) {
				slog.Info("Found matching pre-signed transaction", "tx", preSignedTx)
				break TxInLoop
			}
		}
		slog.Warn("Unknown transaction spending armored outpoint, armor activated", "outpoint", armoredOutPoint)
		actualFee := armoredOutPoint.value
		for _, txOut := range tx.TxOut {
			actualFee -= txOut.Value
		}
		slog.Warn("Actual fee", "fee", actualFee)

		preSignedTx, ok := armoredOutPoint.FindNextPreSignedTx(actualFee)
		if !ok {
			slog.Error("No pre-signed transaction found for fee", "fee", actualFee)
			continue
		}

		res, err := sendTransaction(preSignedTx.raw)
		if err != nil {
			return
		}
		slog.Warn("Sent pre-signed transaction:", "response", res)
	}
}
