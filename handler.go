package main

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

// GetBlockNumberHandler handles GET /block-number
func GetBlockNumberHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"id":      1,
	}

	responseBytes, err := doRPCRequest(requestBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch block number: %v", err), http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, responseBytes)
}

// GetBlockByNumberHandler handles GET /block/{number}
// The number is expected in decimal and will be converted to hex format
func GetBlockByNumberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numberStr := vars["number"]

	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid block number", http.StatusBadRequest)
		return
	}

	hexNumber := fmt.Sprintf("0x%x", number)

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  []interface{}{hexNumber, true},
		"id":      2,
	}

	responseBytes, err := doRPCRequest(requestBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch block data: %v", err), http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, responseBytes)
}

// writeJSONResponse sets the content type and writes response body
func writeJSONResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}