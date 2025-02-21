package internal

import (
	"fmt"
	"github.com/pebbe/zmq4"
	"log"
)

func RunZMQSubscriber(endpoint string) {
	zmqSocket, err := zmq4.NewSocket(zmq4.SUB)
	if err != nil {
		log.Fatal(err)
	}
	defer zmqSocket.Close()

	err = zmqSocket.SetRcvhwm(0)
	if err != nil {
		log.Fatal(err)
	}

	err = zmqSocket.Connect(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to raw transactions
	err = zmqSocket.SetSubscribe("rawtx")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening for Bitcoin Core raw transactions...")

	for {
		// Receive topic
		topic, err := zmqSocket.Recv(0)
		if err != nil {
			log.Fatal(err)
		}
		// Receive message (serialized transaction)
		txBytes, err := zmqSocket.RecvBytes(0)
		if err != nil {
			log.Fatal(err)
		}
		// Receive sequence number
		seqBytes, err := zmqSocket.RecvBytes(0)
		if err != nil {
			log.Fatal(err)
		}
		go handleRawTx(topic, txBytes, seqBytes)
	}
}
