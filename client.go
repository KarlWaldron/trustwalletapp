package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Polygon RPC endpoint
var rpcURL = "https://polygon-rpc.com/"

// doRPCRequest sends a JSON-RPC request to the configured RPC endpoint
func doRPCRequest(body map[string]interface{}) ([]byte, error) {
	jsonPayload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	resp, err := http.Post(rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to send POST request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBytes, nil
}