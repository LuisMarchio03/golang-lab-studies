package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	FindUserInfosUC "github.com/LuisMarchio03/nutri/internal/usecase/user/find"
	updateDietUC "github.com/LuisMarchio03/nutri/internal/usecase/user/update"
)

type UserHandlers struct {
	UpdateDietUsecase    *updateDietUC.UpdateDietUsecase
	FindUserInfosUsecase *FindUserInfosUC.FindUserInfosUsecase
}

func NewUserHandlers(updateDietUsecase *updateDietUC.UpdateDietUsecase, findUserInfosUsecase *FindUserInfosUC.FindUserInfosUsecase) *UserHandlers {
	return &UserHandlers{
		UpdateDietUsecase:    updateDietUsecase,
		FindUserInfosUsecase: findUserInfosUsecase,
	}
}

func (h *UserHandlers) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input updateDietUC.UpdateDietInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(input)
	output, err := h.UpdateDietUsecase.Execute(r.URL.Query().Get("id"), input)
	fmt.Println(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	output.ID = r.URL.Query().Get("id")
	json.NewEncoder(w).Encode(output)
}

func (h *UserHandlers) ListUserHandler(w http.ResponseWriter, r *http.Request) {
	output, err := h.FindUserInfosUsecase.Execute(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}
