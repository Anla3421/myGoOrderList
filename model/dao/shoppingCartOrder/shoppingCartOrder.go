package shoppingCartOrder

import (
	"fmt"
	"server/model/dao"
	"server/model/dto"
)

func ShoppingCartOrder(req *dto.ShoppingCartOrder) error {
	results, err := dao.MysqlConn.Query("INSERT INTO orderGeneral (totalPrice,buyerID) VALUES(?,?) ON DUPLICATE KEY UPDATE totalPrice=?,buyerID=?",
		req.TotalPrice, req.BuyerID, req.TotalPrice, req.BuyerID)
	if err != nil {
		fmt.Println(results)
		return err
	}
	getOrderID, err := dao.MysqlConn.Query("SELECT orderID FROM orderGeneral WHERE buyerID=? ORDER BY orderId DESC LIMIT 0 , 1", req.BuyerID)
	if err != nil {
		fmt.Println(getOrderID)
		return err
	}
	for getOrderID.Next() {
		err = getOrderID.Scan(&req.OrderID)
		if err != nil {
			return err
		}
	}

	for k := range req.ItemID {
		insOrder, err := dao.MysqlConn.Query("INSERT INTO orderDetail (orderID,itemID,amount,price,buyerID) VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE ID=?,itemID=?,amount=?,price=?",
			req.OrderID, req.ItemID[k], req.Amount[k], req.Price[k], req.BuyerID, req.ID, req.ItemID[k], req.Amount[k], req.Price[k])
		if err != nil {
			fmt.Println(insOrder)
			return err
		}
	}

	defer results.Close()
	return nil
}
