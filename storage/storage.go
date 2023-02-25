package storage

import "app/models"

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
	Create(*models.CreateUser) (string, error)
	GetUserById(req *models.UserPrimaryKey) (models.User, error)
	GetList(req *models.GetListRequest) (*models.GetListResponse, error)
	UpdateUser(req *models.UpdateUser) (models.User, error)
	DeleteUser(req *models.UserPrimaryKey) (models.User, error)
}

type ProductRepoI interface {
	CreateProduct(req *models.CreateProduct) (id string, err error)
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
	Create(*models.CreateCategory) (string, error)
	GetByID(*models.CategoryPrimaryKey) (models.Category, error)
	GetAll(*models.GetListCategoryRequest) (models.GetListCategoryResponse, error)
	Update(*models.UpdateCategory, string) error
	Delete(*models.CategoryPrimaryKey) error
}

type BranchRepoI interface {
	CreateBranch(req *models.BranchReq) (id string, err error)
	GetList(req *models.GetBranchListRequest) (*models.GetBranchListResponse, error)
	GetBranchById(req *models.BranchPrimaryKey) (models.Branch, error)
	UpdateBranch(req *models.Branch) (models.Branch, error)
	DeleteBranch(req *models.BranchPrimaryKey) (models.Branch, error)
}
