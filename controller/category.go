package controller

import (
	"net/http"
	"strings"
)

func (c *Controller) Category(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.store.Category().Create(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.store.Category().GetByID(w, r)
		} else {
			c.store.Category().GetAll(w, r)
		}
	}
	if r.Method == "PUT" {
		c.store.Category().Update(w, r)
	}
	if r.Method == "DELETE" {
		c.store.Category().Delete(w, r)
	}
}
