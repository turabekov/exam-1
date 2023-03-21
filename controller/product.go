package controller

import (
	"net/http"
	"strings"
)

func (c *Controller) Product(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.store.Product().CreateProduct(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.store.Product().GetProductById(w, r)
		} else {
			c.store.Product().GetListProduct(w, r)
		}
	}
	if r.Method == "PUT" {
		c.store.Product().UpdateProduct(w, r)
	}
	if r.Method == "DELETE" {
		c.store.Product().DeleteProduct(w, r)
	}
}