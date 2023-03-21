package jsondb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewProductRepo(fileName string, file *os.File) *productRepo {
	return &productRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *productRepo) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// read file
	var products []*models.Product
	err := json.NewDecoder(u.file).Decode(&products)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	var product models.Product
	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	id := uuid.NewString()
	products = append(products, &models.Product{
		Id:         id,
		Name:       product.Name,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	})

	body, err = json.MarshalIndent(products, "", "   ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created successfully!"))
}

// Get list of Products
func (u *productRepo) GetListProduct(w http.ResponseWriter, r *http.Request) {
	// read users from  file
	products := make([]models.Product, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	// get query params limit offset
	var (
		limit    int
		offset   int
		response *models.GetListProductResponse
		e        error
	)

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr == "" {
		limit = len(products)
	} else {
		limit, e = strconv.Atoi(limitStr)
		if e != nil {
			log.Println("strconv err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, e = strconv.Atoi(offsetStr)
		if e != nil {
			log.Println("strconv err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
	}

	if limit+offset > len(products) {
		if offset > len(products) {
			response = &models.GetListProductResponse{
				Count:    len(products),
				Products: []models.Product{},
			}
		} else {
			response = &models.GetListProductResponse{
				Count:    len(products),
				Products: products[offset:],
			}
		}

	} else {
		response = &models.GetListProductResponse{
			Count:    len(products),
			Products: products[offset : limit+offset],
		}
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(body)
}

// Get list by id
func (u *productRepo) GetProductById(w http.ResponseWriter, r *http.Request) {
	// read and unmarshal file
	products := make([]models.Product, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//  read request body
	id := r.URL.Path[len("/product/"):]

	for _, v := range products {
		if v.Id == id {
			body, err := json.Marshal(v)
			if err != nil {
				log.Println("Unmarshal err:", err)
				w.WriteHeader(500)
				w.Write([]byte("Incorrect data"))
				return
			}
			w.WriteHeader(http.StatusFound)
			w.Write(body)
			return
		}
	}

	res := "product with id" + " " + id + " " + "not found"
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(res))
}

// Update product by id
// models.UpdateProduct
func (u *productRepo) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	products := make([]models.Product, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Println("unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	updatedProduct := models.UpdateProduct{}
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}
	//  unmarshal data
	err = json.Unmarshal(body, &updatedProduct)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, v := range products {
		if v.Id == updatedProduct.Id {
			if len(updatedProduct.Name) != 0 {
				products[i].Name = updatedProduct.Name
			}
			if updatedProduct.Price != 0 {
				products[i].Price = updatedProduct.Price
			}
			if len(updatedProduct.CategoryId) != 0 {
				products[i].CategoryId = updatedProduct.CategoryId
			}
			flag = true
		}
	}

	if !flag {
		res := "product with id" + " " + updatedProduct.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}

	body, err = json.MarshalIndent(products, "", "   ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Product updated successfully!"))
}

// Delete user by id
func (u *productRepo) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	products := make([]models.Product, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	var productId models.ProductPrimaryKey
	err = json.Unmarshal(body, &productId)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, v := range products {
		if v.Id == productId.Id {
			products = append(products[:i], products[i+1:]...)
			flag = true
		}
	}

	if !flag {
		res := "product with id" + " " + productId.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}

	body, err = json.MarshalIndent(products, "", "   ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted successfully!"))
}
