package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	getbalance "github.com/marcelofelixsalgado/fullcycle-eda/balances/internal/usecase/get_balance"
)

type WebBalanceHandler struct {
	GetBalanceUseCase getbalance.GetBalanceUseCase
}

func NewWebBalanceHandler(getBalanceUseCase getbalance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUseCase: getBalanceUseCase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "id")

	fmt.Println(accountId)

	dto := getbalance.GetBalanceInputDTO{
		AccountId: accountId,
	}

	output, err := h.GetBalanceUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
