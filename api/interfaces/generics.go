package interfaces

import "WalletViewAPI/api/structure"

type Channels interface {
	structure.UsdTokenPrice | structure.Account
}
