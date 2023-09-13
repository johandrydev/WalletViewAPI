package interfaces

import "net/http"

type AccountHandler interface {
	GetAccountBalance(w http.ResponseWriter, r *http.Request)
}
