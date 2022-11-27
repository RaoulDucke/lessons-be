package handler

import (
	"fmt"

	"github.com/RaolDucke/lessons-be/db"
)

type Handler struct {
	r *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{r: repository}
}

func (h *Handler) AddProduct(p *Product) error {
	err := h.r.AddProduct(convertToDBProduct(p))
	if err != nil {
		return fmt.Errorf("add product^ %s", err)
	}
	return nil
}

func (h *Handler) GetProducts() []*Product {
	products := h.r.GetProducts()
	return convertToProducts(products)
}

func (h *Handler) GetProduct(id int64) (*Product, bool) {
	product, ok := h.r.GetProduct(id)
	if ok {
		return convertToProduct(product), true
	}
	return nil, false
}

func convertToProduct(p *db.Product) *Product {
	return &Product{
		Identity: p.ID,
		Name:     p.Title,
	}
}

func convertToDBProduct(p *Product) *db.Product {
	return &db.Product{
		Title: p.Name,
	}
}

func convertToProducts(products []*db.Product) []*Product {
	res := make([]*Product, 0, len(products))
	for _, p := range products {
		res = append(res, convertToProduct(p))
	}
	return res

}
