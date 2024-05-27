package web

import (
	"encoding/json"
	"net/http"

	"github.com/caiofsr/walletcore/internal/usecases"
)

type WebClientHandler struct {
	CreateClientUseCase usecases.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase usecases.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (wch *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto usecases.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := wch.CreateClientUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
