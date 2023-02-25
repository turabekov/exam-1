package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type Store struct {
	user      *userRepo
	product   *productRepo
	shopCart  *shopCartRepo
	komissiya *komissiyaRepo
	category  *categoryRepo
	branch    *branchRepo
}

func NewFileJson(cfg *config.Config) (storage.StorageI, error) {

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}

	productFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	shopCartFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}

	komissiyaFile, err := os.Open(cfg.Path + cfg.KomissiyaFileName)
	if err != nil {
		return nil, err
	}

	categoryFile, err := os.Open(cfg.Path + cfg.CategoryFileName)
	if err != nil {
		return nil, err
	}

	branchFile, err := os.Open(cfg.Path + cfg.BranchFileName)
	if err != nil {
		return nil, err
	}

	return &Store{
		user:      NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		product:   NewProductRepo(cfg.Path+cfg.ProductFileName, productFile),
		shopCart:  NewShopCartRepo(cfg.Path+cfg.ShopCartFileName, shopCartFile),
		komissiya: NewKomissiyaRepo(cfg.Path+cfg.KomissiyaFileName, komissiyaFile),
		category:  NewCategoryRepo(cfg.Path+cfg.CategoryFileName, categoryFile),
		branch:    NewBranchRepo(cfg.Path+cfg.BranchFileName, branchFile),
	}, nil
}

func (s *Store) CloseDB() {
	s.user.file.Close()
	s.product.file.Close()
	s.shopCart.file.Close()
	s.komissiya.file.Close()
	s.category.file.Close()
	s.branch.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Product() storage.ProductRepoI {
	return s.product
}

func (s *Store) ShopCart() storage.ShopCartRepoI {
	return s.shopCart
}

func (s *Store) Komissiya() storage.KomissiyaRepoI {
	return s.komissiya
}
func (s *Store) Category() storage.CategoryRepoI {
	return s.category
}

func (s *Store) Branch() storage.BranchRepoI {
	return s.branch
}
