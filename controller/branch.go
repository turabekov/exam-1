package controller

import (
	"net/http"
	"strings"
)

func (c *Controller) Branch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.store.Branch().CreateBranch(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.store.Branch().GetBranchById(w, r)
		} else {
			c.store.Branch().GetAll(w, r)
		}
	}
	if r.Method == "PUT" {
		c.store.Branch().UpdateBranch(w, r)
	}
	if r.Method == "DELETE" {
		c.store.Branch().DeleteBranch(w, r)
	}
}