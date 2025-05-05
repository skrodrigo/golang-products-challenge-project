package dtos

type AddFavoriteProductsDTO struct {
	CustomerID  string   `json:"customer_id"`
	ProductsIds []string `json:"products_ids"`
}