package structure

type Asset struct {
	Blockchain        string `json:"blockchain"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimals     int    `json:"tokenDecimals"`
	TokenType         string `json:"tokenType"`
	HolderAddress     string `json:"holderAddress"`
	Balance           string `json:"balance"`
	BalanceRawInteger string `json:"balanceRawInteger"`
	BalanceUsd        string `json:"balanceUsd"`
	TokenPrice        string `json:"tokenPrice"`
	Thumbnail         string `json:"thumbnail"`
	ContractAddress   string `json:"contractAddress,omitempty"`
}

type Account struct {
	TotalBalanceUsd string  `json:"totalBalanceUsd"`
	TotalCount      int     `json:"totalCount"`
	Assets          []Asset `json:"assets"`
}

type UsdTokenPrice struct {
	UsdPrice        string `json:"usdPrice"`
	Blockchain      string `json:"blockchain"`
	ContractAddress string `json:"contractAddress"`
}
