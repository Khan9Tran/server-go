package user

import (
	"github.com/go-playground/validator/v10"
	"ecom/service/auth"
	"ecom/types"
	"ecom/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
type Handler struct {
	store types.UserStore
}

func NewHandler(Store types.UserStore) *Handler { 
	return &Handler{store: Store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with %s already exists", payload.Email))
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}


	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}

