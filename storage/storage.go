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
	GetListProduct(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	GetProductById(req *models.ProductPrimaryKey) (models.Product, error)
	UpdateProduct(req *models.UpdateProduct) (models.Product, error)
	DeleteProduct(req *models.ProductPrimaryKey) (models.Product, error)
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
	CreateBranch(req *models.BranchReq) (id string, err error)
	GetList(req *models.GetBranchListRequest) (*models.GetBranchListResponse, error)
	GetBranchById(req *models.BranchPrimaryKey) (models.Branch, error)
	UpdateBranch(req *models.Branch) (models.Branch, error)
	DeleteBranch(req *models.BranchPrimaryKey) (models.Branch, error)
}
