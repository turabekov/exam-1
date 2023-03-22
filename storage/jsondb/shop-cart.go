package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

type shopCartRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewShopCartRepo(fileName string, file *os.File) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (s *shopCartRepo) AddShopCart(req *models.AddShopCart) (models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return models.ShopCart{}, err
	}

	// if userId and productId exist replace only count
	newShopCart := models.ShopCart{}

	now := time.Now()

	newShopCart = models.ShopCart{
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Count:     req.Count,
		Status:    false,
		Time:      now.Format("2006-01-02 15:04:05"),
	}
	carts = append(carts, newShopCart)

	// stringify struct to json
	body, err := json.MarshalIndent(carts, "", "   ")
	if err != nil {
		return models.ShopCart{}, err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return models.ShopCart{}, err
	}

	return newShopCart, nil
}

func (s *shopCartRepo) RemoveShopCart(req *models.RemoveShopCart) (models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return models.ShopCart{}, err
	}

	deletedShopCart := models.ShopCart{}
	flag := false
	for i, v := range carts {
		if v.UserId == req.UserId && v.ProductId == req.ProductId {
			deletedShopCart = carts[i]
			carts = append(carts[:i], carts[i+1:]...)
			flag =  true
		}
	}

	if !flag {
		return models.ShopCart{}, errors.New("shop-cart not found")
	}

	// stringify struct to json
	body, err := json.MarshalIndent(carts, "", "   ")
	if err != nil {
		return models.ShopCart{}, err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return models.ShopCart{}, err
	}

	return deletedShopCart, nil
}

func (s *shopCartRepo) GetAllShopCarts() ([]models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return []models.ShopCart{}, err
	}

	// Task 12
	// sorting desc by default
	sort.Slice(carts, func(i, j int) bool {
		firstDate, err := time.Parse("2006-01-02 15:04:05", carts[i].Time)
		if err != nil {
			return false
		}
		secondDate, err := time.Parse("2006-01-02 15:04:05", carts[j].Time)
		if err != nil {
			return false
		}
		return firstDate.After(secondDate)
	})

	return carts, nil
}

func (s *shopCartRepo) GetUserShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return []models.ShopCart{}, err
	}

	userShopCarts := []models.ShopCart{}

	for _, v := range carts {
		if v.UserId == req.Id {
			userShopCarts = append(userShopCarts, v)
		}
	}

	// Task 12
	// sorting desc by default
	sort.Slice(userShopCarts, func(i, j int) bool {
		firstDate, err := time.Parse("2006-01-02 15:04:05", userShopCarts[i].Time)
		if err != nil {
			return false
		}
		secondDate, err := time.Parse("2006-01-02 15:04:05", userShopCarts[j].Time)
		if err != nil {
			return false
		}
		return firstDate.After(secondDate)
	})

	return userShopCarts, nil
}

func (s *shopCartRepo) UpdateShopCart(cart models.ShopCart) (models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return models.ShopCart{}, err
	}

	updatedShopCart := models.ShopCart{}
	for i, v := range carts {
		if v.ProductId == cart.ProductId && v.UserId == cart.UserId {
			carts[i].ProductId = cart.ProductId
			carts[i].UserId = cart.UserId
			carts[i].Status = cart.Status

			updatedShopCart = carts[i]
		}
	}

	if len(updatedShopCart.ProductId) <= 0 {
		return models.ShopCart{}, errors.New("shop-cart not found")
	}

	// stringify struct to json
	body, err := json.MarshalIndent(carts, "", "   ")
	if err != nil {
		return models.ShopCart{}, err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return models.ShopCart{}, err
	}

	return updatedShopCart, nil
}
