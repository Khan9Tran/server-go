package product

import (
	"database/sql"
	"ecom/types"
)

type Store struct{
	db *sql.DB
}

func NewProduct(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error){
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next(){
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error){
	p := new(types.Product)
	err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt, &p.Quantity, &p.Description, &p.Image)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *Store)CreateProduct(p types.Product) error{
	_, err := s.db.Exec("INSERT INTO products(name, price, description, quantity, image) VALUES(?, ?, ?, ?, ?)", p.Name, p.Price, p.Description, p.Quantity, p.Image)
	if err != nil {
		return err
	}
	return nil
}