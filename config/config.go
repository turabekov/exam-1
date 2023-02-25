package config

type Config struct {
	Path string

	UserFileName      string
	ProductFileName   string
	ShopCartFileName  string
	KomissiyaFileName string
	CategoryFileName  string
	BranchFileName    string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/user.json"
	cfg.ProductFileName = "/product.json"
	cfg.ShopCartFileName = "/shop_cart.json"
	cfg.KomissiyaFileName = "/comission.json"
	cfg.CategoryFileName = "/category.json"
	cfg.BranchFileName = "/branch.json"

	return cfg
}
