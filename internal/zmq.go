package internal

import (
	"github.com/pebbe/zmq4"
	"log/slog"
	"os"
)

func RunZMQSubscriber(endpoint string) {
	zmqSocket, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		slog.Error("Failed to create ZMQ socket", "err", err)
		os.Exit(1)
	}
	defer zmqSocket.Close()

	err = zmqSocket.SetRcvhwm(0)
	if err != nil {
		slog.Error("Failed to set ZMQ RCVHWM", "err", err)
		os.Exit(1)
	}

	err = zmqSocket.Connect(endpoint)
	if err != nil {
		slog.Error("Failed to connect to ZMQ endpoint", "err", err)
		os.Exit(1)
	}

	// Subscribe to raw transactions
	err = zmqSocket.SetSubscribe("rawtx")
	if err != nil {
		slog.Error("Failed to subscribe to raw transactions", "err", err)
		os.Exit(1)
	}

	slog.Info("Listening for Bitcoin Core raw transactions...")

	for {
		// Receive topic
		topic, err := zmqSocket.Recv(0)
		if err != nil {
			slog.Error("Failed to receive topic", "err", err)
		}
		// Receive message (serialized transaction)
		txBytes, err := zmqSocket.RecvBytes(0)
		if err != nil {
			slog.Error("Failed to receive transaction", "err", err)
		}
		// Receive sequence number
		seqBytes, err := zmqSocket.RecvBytes(0)
		if err != nil {
			slog.Error("Failed to receive sequence number", "err", err)
		}
		go handleRawTx(topic, txBytes, seqBytes)
	}
}
