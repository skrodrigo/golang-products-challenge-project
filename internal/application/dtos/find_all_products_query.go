package dtos

type FindAllProductsQuery struct {
	NameLike        string  `json:"name_like"`
	DescriptionLike string  `json:"description_like"`
	PriceMin        float64 `json:"price_min"`
	PriceMax        float64 `json:"price_max"`
	Active          bool    `json:"active"`
}