package shoppingCartUpdate

import (
	"fmt"
	"server/model/dao"
	"server/model/dto"
)

func ShoppingCartUpdate(req *dto.ShoppingCartUpdate) error {
	tx, err := dao.MysqlConn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for k := range req.ItemID {
		if _, err = tx.Exec("UPDATE orderDetail SET itemID=?,amount=?,price=? WHERE ID=?",
			req.ItemID[k], req.Amount[k], req.Price[k], req.ID[k]); err != nil {
			tx.Rollback()
			fmt.Println("update failed:", err)
			return err
		}

	}
	_, err = dao.MysqlConn.Exec("UPDATE orderGeneral SET totalPrice=? WHERE orderID=?",
		req.TotalPrice, req.OrderID)
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("update success")
	return tx.Commit()
}
