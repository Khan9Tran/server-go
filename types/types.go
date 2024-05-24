package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(Product) error
}

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"createdAt"`
	Quantity int `json:"quantity"`
	Image string `json:"image"`
}

type User struct{
	ID int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
type LoginUserPayload struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateProductPayload struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price float64 `json:"price" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
	Image string `json:"image" validate:"required"`
}