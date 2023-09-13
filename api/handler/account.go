package handler

import (
	"WalletViewAPI/api/interfaces"
	"WalletViewAPI/api/response"
	"WalletViewAPI/api/service/ankr"
	"WalletViewAPI/api/structure"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Account struct {
	service interfaces.BalanceProvider
}

func NewAccount() interfaces.AccountHandler {
	return &Account{
		service: ankr.NewAccount(),
	}
}

type ch[T interfaces.Channels] struct {
	data T
	err  error
}

func (account Account) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		response.Create(w, "address is required", http.StatusBadRequest, nil)
		return
	}
	currency := r.URL.Query().Get("currency")
	if currency == "" {
		response.Create(w, "currency is required", http.StatusBadRequest, nil)
		return
	}

	tknCh := make(chan ch[structure.UsdTokenPrice])
	go func() {
		CurrencyUsdTokenPrice, err := account.service.GetTokenPrice(strings.ToLower(currency), "")
		if err != nil {
			log.Printf("error: %v\n", err)
			tknCh <- ch[structure.UsdTokenPrice]{err: err}
		}
		tknCh <- ch[structure.UsdTokenPrice]{data: CurrencyUsdTokenPrice}
	}()

	balanceCh := make(chan ch[structure.Account])
	go func() {
		balance, err := account.service.GetBalance(address)
		if err != nil {
			log.Printf("error: %v\n", err)
			balanceCh <- ch[structure.Account]{err: err}
		}
		balanceCh <- ch[structure.Account]{data: balance}
	}()

	tkn, balance := <-tknCh, <-balanceCh
	if tkn.err != nil || balance.err != nil {
		response.Create(w, "error", http.StatusInternalServerError, nil)
		return
	}

	tokenCurrencyPrice, err := strconv.ParseFloat(tkn.data.UsdPrice, 64)
	if err != nil {
		log.Printf("error: %v\n", err)
		response.Create(w, "error", http.StatusInternalServerError, nil)
		return
	}

	balances, err := account.service.GetBalanceExchangeRates(balance.data.Assets, tokenCurrencyPrice)
	if err != nil {
		log.Printf("error: %v\n", err)
		response.Create(w, "error", http.StatusInternalServerError, nil)
		return
	}

	response.Create(w, "success", http.StatusOK, balances)
}
