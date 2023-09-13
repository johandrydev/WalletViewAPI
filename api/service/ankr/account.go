package ankr

import (
	"WalletViewAPI/api/interfaces"
	"WalletViewAPI/api/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Account struct{}

func NewAccount() interfaces.BalanceProvider {
	return &Account{}
}

const urlApi = "https://rpc.ankr.com/multichain/79258ce7f7ee046decc3b5292a24eb4bf7c910d7e39b691384c7ce0cfb839a01/?"

func (Account) GetBalance(address string) (structure.Account, error) {
	url := urlApi + "ankr_getAccountBalance="
	if address == "" {
		return structure.Account{}, fmt.Errorf("address is required")
	}

	bodyReq := structure.AnkrReqBody{
		Jsonrpc: "2.0",
		Method:  "ankr_getAccountBalance",
		Params: structure.AnkrGetAccountReq{
			WalletAddress: address,
		},
		ID: 1,
	}
	bodyJson, err := json.Marshal(bodyReq)
	if err != nil {
		return structure.Account{}, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(bodyJson)))
	if err != nil {
		return structure.Account{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return structure.Account{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return structure.Account{}, err
	}

	var ankrResponse structure.AnkrGetBalanceResponse
	err = json.Unmarshal(body, &ankrResponse)
	if err != nil {
		return structure.Account{}, err
	}
	return ankrResponse.Result, nil
}

func (Account) GetTokenPrice(blockchain string, contractAddress string) (structure.UsdTokenPrice, error) {
	url := urlApi + "ankr_getTokenPrice="
	if blockchain == "" {
		return structure.UsdTokenPrice{}, fmt.Errorf("blockchain is required")
	}
	bodyReq := structure.AnkrReqBody{
		Jsonrpc: "2.0",
		Method:  "ankr_getTokenPrice",
		Params: structure.AnkrTokenPriceReq{
			Blockchain:      blockchain,
			ContractAddress: contractAddress,
		},
		ID: 1,
	}
	bodyJson, err := json.Marshal(bodyReq)
	if err != nil {
		return structure.UsdTokenPrice{}, err
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(bodyJson)))
	if err != nil {
		return structure.UsdTokenPrice{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return structure.UsdTokenPrice{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return structure.UsdTokenPrice{}, err
	}

	var ankrResponse structure.AnkrGetTokenPriceResponse
	err = json.Unmarshal(body, &ankrResponse)
	if err != nil {
		return structure.UsdTokenPrice{}, err
	}
	return ankrResponse.Result, nil
}

func (Account) GetBalanceExchangeRates(assets []structure.Asset, tknCurrencyPrice float64) (map[string]string, error) {
	balances := make(map[string]string)
	for _, asset := range assets {
		BalanceUsd, err := strconv.ParseFloat(asset.BalanceUsd, 64)
		if err != nil {
			return nil, err
		}

		tokenPrice, err := strconv.ParseFloat(asset.TokenPrice, 64)
		if err != nil {
			return nil, err
		}

		exchangeRate := tokenPrice / tknCurrencyPrice
		BalanceCurrency := BalanceUsd * exchangeRate
		balances[asset.TokenSymbol] = strconv.FormatFloat(BalanceCurrency, 'f', 8, 64)
	}
	return balances, nil
}
