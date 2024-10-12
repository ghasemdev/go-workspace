package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "Amir Azimi",
	"type": "deposit",
	"amount" : 10000000.5
}
`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	// Simulate a file/socket
	reader := bytes.NewBufferString(data)

	// Decode the request
	decode := json.NewDecoder(reader)

	var request Request

	if err := decode.Decode(&request); err != nil {
		log.Fatalf("Error: Can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", request)

	prevBalance := 8500000.0
	response := map[string]interface{}{
		"ok":      true,
		"code":    200,
		"balance": prevBalance + request.Amount,
	}

	encode := json.NewEncoder(os.Stdout)
	if err := encode.Encode(response); err != nil {
		log.Fatalf("Error: Can't encode - %s", err)
	}
}
