package common

type Order struct {
	UserID        int32   `json:"user_id"`
	UniqueOrderID string  `json:"unique_order_id"`
	TotalPrice    float64 `json:"total_price"`
}

type SubOrder struct {
	UniqueOrderID string `json:"unique_order_id"`
	ProductID     int32  `json:"product_id"`
	Quantity      int32  `json:"quantity"`
	Status        string `json:"status"`
}

type GetOrderRequest struct {
	UniqueOrderID string `json:"unique_order_id"`
	UserID        int32  `json:"user_id"`
}

type GetOrderResponse struct {
	UniqueOrderID    string              `json:"unique_order_id"`
	UserID           int32               `json:"user_id"`
	TotalPrice       float64             `json:"total_price"`
	Orderstatus      string              `json:"order_status"`
	SubOrderResponse []*SubOrderResponse `json:"data"`
}

type SubOrderResponse struct {
	ProductID      int32  `json:"product_id"`
	Quantity       int32  `json:"quantity"`
	Suborderstatus string `json:"status"`
}
