package db

import (
	"database/sql"
	"errors"

	"golang.org/x/net/context"
)

type Repository struct {
	products []*Product
	users    []*User

	database *sql.DB
}

func New(database *sql.DB) *Repository {
	return &Repository{
		products: []*Product{},
		users:    []*User{},

		database: database,
	}

}
func (r *Repository) AddProduct(ctx context.Context, p *Product) error {
	if p == nil {
		return errors.New("product is nil")
	}
	if p.Title == "" {
		return errors.New("title is empty")
	}
	if p.Price <= 0 {
		return errors.New("price <=0")
	}
	id := int64(1)
	if len(r.products) > 0 {
		lastProduct := r.products[len(r.products)-1]
		id = lastProduct.ID + 1
	}

	p.ID = id
	r.products = append(r.products, p)
	return nil

	// 	_, err := r.database.ExecContext(ctx, `
	// 		insert into product (title, price)
	// 		values ($1,$2)
	// 	`, p.Title, p.Price)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
}

func (r *Repository) UpdateProduct(ctx context.Context, p *Product) (bool, error) {
	if p == nil {
		return false, errors.New("product is nil")
	}
	if p.Title == "" {
		return false, errors.New("title is empty")
	}
	if p.ID <= 0 {
		return false, errors.New("id <= 0")
	}
	product, ok := r.GetProduct(p.ID)
	if ok {
		if product.Title != "" {
			product.Title = p.Title
		}
		if p.Price > 0 {
			product.Price = p.Price
		}
	}
	return ok, nil
}

func (r *Repository) GetProducts(ctx context.Context) ([]*Product, error) {
	raws, err := r.database.QueryContext(ctx, `
	select id,title, price
	from product
	`)
	if err != nil {
		return nil, err
	}

	defer raws.Close()
	var result []*Product

	for raws.Next() {
		p := new(Product)
		err = raws.Scan(&p.ID, &p.Title, &p.Price)
		if err != nil {
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}

func (r *Repository) GetProduct(id int64) (*Product, bool) {
	for _, product := range r.products {
		if id == product.ID {
			return product, true
		}

	}
	return nil, false
}

func (r *Repository) DoesProductExist(id int64) bool {
	for _, product := range r.products {
		if id == product.ID {
			return true
		}

	}
	return false
}

func (r *Repository) AddUser(p *User) error {
	if p == nil {
		return errors.New("user is nil")
	}
	if p.Name == "" {
		return errors.New("name is empty")
	}
	if p.Surname == "" {
		return errors.New("surname is empty")
	}
	id := int64(1)
	if len(r.users) > 0 {
		lastUser := r.users[len(r.users)-1]
		id = lastUser.Id + 1
	}

	p.Id = id
	r.users = append(r.users, p)
	return nil
}

func (r *Repository) GetUsers() []*User {
	return r.users
}
func (r *Repository) GetUser(id int64) (*User, bool) {
	for _, user := range r.users {
		if id == user.Id {
			return user, true
		}

	}
	return nil, false
}
