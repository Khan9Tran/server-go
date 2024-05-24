package product

import (
	"ecom/types"
	"net/http"
	"ecom/utils"
	"github.com/gorilla/mux"
	"github.com/go-playground/validator/v10"
	"fmt"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler{
	return &Handler{store: store}
}

func (h *Handler) RegisterRouter(router *mux.Router){
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodGet)
	router.HandleFunc("/addProduct", h.handleAddProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request){
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) handleAddProduct(w http.ResponseWriter, r *http.Request){
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateProduct(types.Product{
		Name: payload.Name,
		Price: payload.Price,
		Quantity: payload.Quantity,
		Image: payload.Image,
		Description: payload.Description,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)


}