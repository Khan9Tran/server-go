package user

import (
	"bytes"
	"ecom/types"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T){
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "valida@mail.com",
			Password:  "password",
		}
		marshalled, _ := json.Marshal(payload)
		
	
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
	
		rr := httptest.NewRecorder()
		route := mux.NewRouter()
	
		route.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)
		route.ServeHTTP(rr, req)
	
		if rr.Code != http.StatusCreated{
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct {}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error){
	return nil, nil
}
func (m *mockUserStore) GetUserByID(id int) (*types.User, error){
	return nil, nil
}
func (m *mockUserStore) CreateUser(u types.User) error{
	return nil
}