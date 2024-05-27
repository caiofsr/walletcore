package web

import (
	"encoding/json"
	"net/http"

	"github.com/caiofsr/walletcore/internal/usecases"
)

type WebTransferHandler struct {
	CreateTransferUseCase usecases.CreateTransferUseCase
}

func NewWebTransferHandler(createTransferUseCase usecases.CreateTransferUseCase) *WebTransferHandler {
	return &WebTransferHandler{
		CreateTransferUseCase: createTransferUseCase,
	}
}

func (wth *WebTransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var dto usecases.CreateTransferInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := wth.CreateTransferUseCase.Execute(dto)
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
