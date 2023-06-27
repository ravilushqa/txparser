package ethereum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	MaxRetries    = 4
	RetryInterval = 3 * time.Second
)

type RPCRequest struct {
	Jsonrpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      int      `json:"id"`
}

type RPCResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	ID      int             `json:"id"`
}

type Client struct {
	BaseURL string
}

func NewClient(url string) *Client {
	return &Client{
		BaseURL: url,
	}
}

func (c *Client) SendRequest(r RPCRequest) (RPCResponse, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return RPCResponse{}, err
	}

	var resp *http.Response
	for i := 0; i < MaxRetries; i++ {
		resp, err = http.Post(c.BaseURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			// if err, wait before next try
			time.Sleep(RetryInterval)
			continue
		}
	}
	if err != nil {
		return RPCResponse{}, fmt.Errorf("error sending request: %s", err)
	}

	var rpcResponse RPCResponse
	err = json.NewDecoder(resp.Body).Decode(&rpcResponse)
	if err != nil {
		return RPCResponse{}, fmt.Errorf("error decoding response: %s", err)
	}

	return rpcResponse, nil
}
