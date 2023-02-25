package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `json:"category_id"`
}

type Product struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `json:"category_id"`
}

type UpdateProduct struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryId string `json:"category_id"`
}

type GetListProductRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListProductResponse struct {
	Count    int       `json:"count"`
	Products []Product `json:"products"`
}
