package handler

import (
	"encoding/json"
	"gitee.com/geekbang/basic-go/web3/go-backend/db"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	address := strings.TrimPrefix(r.URL.Path, "/transactions/")
	if !common.IsHexAddress(address) {
		http.Error(w, "Invalid address", http.StatusBadRequest)
		return
	}

	transactions, err := db.GetTransactionsByAddress(address)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
