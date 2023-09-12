package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	url := "https://rpc.ankr.com/multichain/79258ce7f7ee046decc3b5292a24eb4bf7c910d7e39b691384c7ce0cfb839a01/?ankr_getAccountBalance="

	payload := strings.NewReader("{\"jsonrpc\":\"2.0\",\"method\":\"ankr_getAccountBalance\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
