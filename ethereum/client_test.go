package ethereum

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSendRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"jsonrpc": "2.0", "result": "OK", "id": 1}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	client := NewClient(ts.URL)
	request := RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []string{},
		ID:      1,
	}

	resp, err := client.SendRequest(request)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Jsonrpc != "2.0" {
		t.Errorf("Expected Jsonrpc to be '2.0', but got '%s'", resp.Jsonrpc)
	}
	if string(resp.Result) != `"OK"` {
		t.Errorf("Expected Result to be 'OK', but got '%s'", resp.Result)
	}
	if resp.ID != 1 {
		t.Errorf("Expected Id to be 1, but got '%d'", resp.ID)
	}
}

func TestSendRequestWithError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	client := NewClient(ts.URL)
	request := RPCRequest{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []string{},
		ID:      1,
	}

	_, err := client.SendRequest(request)
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}
