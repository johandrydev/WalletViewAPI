package structure

type AnkrGetAccountReq struct {
	WalletAddress string `json:"walletAddress"`
}

type AnkrTokenPriceReq struct {
	Blockchain      string `json:"blockchain"`
	ContractAddress string `json:"contractAddress"`
}

type AnkrReqBody struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      int    `json:"id"`
}

type AnkrResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
}

type AnkrGetBalanceResponse struct {
	AnkrResponse
	Result Account `json:"result"`
}

type AnkrGetTokenPriceResponse struct {
	AnkrResponse
	Result UsdTokenPrice `json:"result"`
}
