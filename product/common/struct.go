package common

type Product struct {
	Id           int32   `json:"id"`
	Name         string  `json:"name"`
	Weight       int32   `json:"weight"` // 500, 1
	Unit         string  `json:"unit"`   // grams, litre
	Quantity     int32   `json:"quantity"`
	PricePerItem float64 `json:"price_per_item"`
	Version      int     `json:"version"`
}

type UpdateProductQuantity struct {
	Id                     int32 `json:"id"`
	DescreaseQuantityCount int32 `json:"quatity_count"`
}
type AuthorizationTokenRequest struct {
	Token string `json:"token"`
}
type AuthorizationTokenResponse struct {
	IsAuthorized bool   `json:"is_authorized"`
	Email        string `json:"email"`
}
