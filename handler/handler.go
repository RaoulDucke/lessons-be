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
	return convertToProducts(h.r.GetProducts())
}

func convertToDBProduct(p *Product) *db.Product {
	return &db.Product{
		Title: p.Title,
	}
}

func convertToProducts(products []*db.Product) []*Product {
	res := make([]*Product, 0, len(products))
	for _, p := range products {
		res = append(res, &Product{Title: p.Title})
	}
	return res

}
