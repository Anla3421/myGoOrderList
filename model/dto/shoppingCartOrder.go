package dto

type ShoppingCartOrder struct {
	TotalPrice int
	BuyerID    int `form:"buyerID" json:"buyerID,omitempty"`

	ItemID []int `form:"itemID" json:"itemID,omitempty"`
	Amount []int `form:"amount" json:"amount,omitempty"`
	Price  []int `form:"price" json:"price,omitempty"`

	TimeNow string `form:"timeNow" json:"timeNow,omitempty"`

	OrderID int
	ID      int
}
