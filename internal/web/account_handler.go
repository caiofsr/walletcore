package web

import (
	"encoding/json"
	"net/http"

	"github.com/caiofsr/walletcore/internal/usecases"
)

type WebAccountHandler struct {
	CreateAccountUseCase usecases.CreateAccountUseCase
}

func NewWebAccountHandler(createAccountUseCase usecases.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		CreateAccountUseCase: createAccountUseCase,
	}
}

func (wah *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var dto usecases.CreateAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := wah.CreateAccountUseCase.Execute(dto)
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
