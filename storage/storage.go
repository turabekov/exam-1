package storage

import (
	"app/models"
	"net/http"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
	Komissiya() KomissiyaRepoI
	Category() CategoryRepoI
	Branch() BranchRepoI
}

type UserRepoI interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetList(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type ProductRepoI interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetListProduct(w http.ResponseWriter, r *http.Request)
	GetProductById(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
}

type ShopCartRepoI interface {
	AddShopCart(req *models.AddShopCart) (models.ShopCart, error)
	RemoveShopCart(req *models.RemoveShopCart) (models.ShopCart, error)
	GetAllShopCarts() ([]models.ShopCart, error)
	GetUserShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error)
	UpdateShopCart(cart models.ShopCart) (models.ShopCart, error)
}

type KomissiyaRepoI interface {
	GetKomissiya() (models.Komissiya, error)
	UpdateBalanceKomissiya(komissiya models.Komissiya) error
}

type CategoryRepoI interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type BranchRepoI interface {
	CreateBranch(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetBranchById(w http.ResponseWriter, r *http.Request)
	UpdateBranch(w http.ResponseWriter, r *http.Request)
	DeleteBranch(w http.ResponseWriter, r *http.Request)
}
