package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type AddShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
}

type ShopCart struct {
	// Id        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type RemoveShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type UserProductIds struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type CalculateShop struct {
	UserID         string `json:"user_id"`
	Discount       int    `json:"discount"`
	DiscountStatus string `json:"discount_status"`
}
