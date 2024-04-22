package common

type Product struct {
	Id       int32
	Name     string
	Weight   int32  // 500, 1
	Unit     string // grams, litre
	Quantity int32
	Version  int
}

type UpdateProductQuantity struct {
	Id                     int32 `json:"id"`
	DescreaseQuantityCount int32 `json:"quatity_count"`
}
