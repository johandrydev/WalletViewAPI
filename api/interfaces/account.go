package interfaces

import "WalletViewAPI/api/structure"

type BalanceProvider interface {
	GetBalance(address string) (structure.Account, error)
	GetTokenPrice(blockchain string, contractAddress string) (structure.UsdTokenPrice, error)
	GetBalanceExchangeRates(assets []structure.Asset, tknCurrencyPrice float64) (map[string]string, error)
}
