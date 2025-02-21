package internal

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const poolUrl = "https://mempool.space/testnet/api/tx"

func sendTransaction(tx []byte) (string, error) {
	resp, err := http.Post(poolUrl, "text/plain", bytes.NewBuffer(tx))
	if err != nil {
		return "", fmt.Errorf("failed to send transaction: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}
