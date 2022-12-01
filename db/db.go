package db

import "errors"

type Repository struct {
	products []*Product
}

func New() *Repository {
	return &Repository{
		products: []*Product{},
	}

}
func (r *Repository) AddProduct(p *Product) error {
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
}

func (r *Repository) UpdateProduct(p *Product) (bool, error) {
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

func (r *Repository) GetProducts() []*Product {
	return r.products
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
