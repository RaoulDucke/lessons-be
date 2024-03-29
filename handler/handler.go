package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/RaolDucke/lessons-be/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	r *db.Repository
}

func New(repository *db.Repository) *Handler {
	return &Handler{r: repository}
}

func (h *Handler) AddProduct(ctx context.Context, c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		internalError(c, err)
		return
	}
	product := new(Product)
	err = json.Unmarshal(jsonData, product)
	if err != nil {
		internalError(c, err)
		return
	}
	if product.Name == "" {
		badRequst(c)
		return
	}
	if product.Price <= 0 {
		badRequst(c)
		return
	}
	err = h.r.AddProduct(ctx, convertToDBProduct(product))
	if err != nil {
		internalError(c, err)
		return
	}
}

func (h *Handler) GetProducts(ctx context.Context, c *gin.Context) {
	idString := c.Request.URL.Query().Get("id")
	if idString != "" {
		id, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			badRequst(c)
			return
		}
		product, ok := h.getProduct(id)
		if ok {
			statusOk(c, product)
		} else {
			notFound(c)
		}
		return
	}
	products, err := h.r.GetProducts(ctx)
	if err != nil {
		internalError(c, err)
		return
	}

	c.JSON(http.StatusOK, convertToProducts(products))

}

func (h *Handler) UpdateProduct(ctx context.Context, c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		internalError(c, err)
		return
	}
	product := new(Product)
	err = json.Unmarshal(jsonData, product)
	if err != nil {
		internalError(c, err)
		return
	}
	if product.Name == "" && product.Price <= 0 {
		badRequst(c)
		return
	}
	if product.Identity <= 0 {
		badRequst(c)
		return
	}
	ok, err := h.r.UpdateProduct(ctx, convertToDBProduct(product))
	if err != nil {
		internalError(c, err)
		return
	}
	if !ok {
		notFound(c)
		return
	}
}

func (h *Handler) getProduct(id int64) (*Product, bool) {
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
		Price:    p.Price,
	}
}

func convertToDBProduct(p *Product) *db.Product {
	return &db.Product{
		Title: p.Name,
		ID:    p.Identity,
		Price: p.Price,
	}
}

func convertToProducts(products []*db.Product) []*Product {
	res := make([]*Product, 0, len(products))
	for _, p := range products {
		res = append(res, convertToProduct(p))
	}
	return res

}

func internalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal error: %s", err))
}

func badRequst(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "bad request")
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, "not found")
}

func statusOk(c *gin.Context, val any) {
	c.JSON(http.StatusOK, val)
}
