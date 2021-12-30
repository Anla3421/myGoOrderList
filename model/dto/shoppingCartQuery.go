package dto

type ShoppingCartQueryRes struct {
	ID         int    `json:"ID,omitempty"`
	OrderID    int    `json:"orderID,omitempty"`
	TotalPrice int    `json:"totalPrice,omitempty"`
	Finish     string `json:"finish,omitempty"`

	Account string   `json:"account,omitempty"`
	Name    []string `json:"name,omitempty"`

	ItemID []int64 `json:"itemID,omitempty"`
	Amount []int64 `json:"amount,omitempty"`
	Price  []int64 `json:"price,omitempty"`
}
