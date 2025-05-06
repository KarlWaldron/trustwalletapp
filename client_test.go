package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestDoRPCRequest_Success tests a successful JSON-RPC response
func TestDoRPCRequest_Success(t *testing.T) {
	// Mock server that returns a fixed block number
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  "0x123456",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer mockServer.Close()

	// Override global rpcURL to use mock server
	originalURL := rpcURL
	rpcURL = mockServer.URL
	defer func() { rpcURL = originalURL }()

	body := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"id":      1,
	}

	respBytes, err := doRPCRequest(body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	var resp map[string]interface{}
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	if resp["result"] != "0x123456" {
		t.Errorf("Expected block number '0x123456', got '%v'", resp["result"])
	}
}

// TestDoRPCRequest_Non200 tests handling of non-200 status code
func TestDoRPCRequest_Non200(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	originalURL := rpcURL
	rpcURL = mockServer.URL
	defer func() { rpcURL = originalURL }()

	body := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"id":      1,
	}

	_, err := doRPCRequest(body)
	if err == nil {
		t.Fatal("Expected error on non-200 response, got nil")
	}
}