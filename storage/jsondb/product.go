package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
func (u *productRepo) GetListProduct(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	if req.Limit+req.Offset > len(products) {
		if req.Offset > len(products) {
			return &models.GetListProductResponse{
				Count:    len(products),
				Products: []models.Product{},
			}, nil
		}

		return &models.GetListProductResponse{
			Count:    len(products),
			Products: products[req.Offset:],
		}, nil
	}

	response := &models.GetListProductResponse{
		Count:    len(products),
		Products: products[req.Offset : req.Limit+req.Offset],
	}

	return response, nil
}

// Get list by id
func (u *productRepo) GetProductById(req *models.ProductPrimaryKey) (models.Product, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.Product{}, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	for _, v := range products {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.Product{}, errors.New("product not found")
}

// Update user by id
func (u *productRepo) UpdateProduct(req *models.UpdateProduct) (models.Product, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.Product{}, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	updatedUser := models.Product{}
	for i, v := range products {
		if v.Id == req.Id {
			if len(req.Name) != 0 {
				products[i].Name = req.Name
			}
			if req.Price != 0 {
				products[i].Price = req.Price
			}
			if len(req.CategoryId) != 0 {
				products[i].CategoryId = req.CategoryId
			}
			updatedUser = products[i]
		}
	}

	if len(updatedUser.Name) <= 0 {
		return models.Product{}, errors.New("product not found")
	}

	body, err := json.MarshalIndent(products, "", "   ")

	if err != nil {
		return models.Product{}, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.Product{}, err
	}

	return updatedUser, nil

}

// Delete user by id
func (u *productRepo) DeleteProduct(req *models.ProductPrimaryKey) (models.Product, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.Product{}, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	deletedUser := models.Product{}
	for i, v := range products {
		if v.Id == req.Id {
			deletedUser = products[i]
			products = append(products[:i], products[i+1:]...)
		}
	}

	if len(deletedUser.Name) <= 0 {
		return models.Product{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(products, "", "   ")

	if err != nil {
		return models.Product{}, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.Product{}, err
	}

	return deletedUser, nil

}
