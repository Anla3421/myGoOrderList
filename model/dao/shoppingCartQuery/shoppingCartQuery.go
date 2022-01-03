package shoppingCartQuery

import (
	"server/model/dao"
	"server/model/dto"
)

func ShoppingCartQuery(BuyerID int, Account string) (res dto.ShoppingCartQueryRes, err error) {
	results, err := dao.MysqlConn.Query("SELECT "+
		"orderGeneral.orderID, orderGeneral.totalPrice, orderGeneral.finish "+
		"FROM `orderGeneral` "+
		" WHERE orderGeneral.buyerID=? and orderGeneral.finish=? ", BuyerID, "no")
	if err != nil {
		return res, err
	}

	for results.Next() {
		err = results.Scan(&res.OrderID, &res.TotalPrice, &res.Finish)
		if err != nil {

			return res, err
		}

	}
	res.Account = Account

	if res.OrderID != 0 {
		results, err := dao.MysqlConn.Query("SELECT "+
			"orderDetail.Amount,orderDetail.Price, "+
			"goods.name "+
			"FROM orderDetail "+
			"LEFT JOIN goods ON orderDetail.itemID = goods.id "+
			"WHERE orderDetail.OrderID=?", res.OrderID)
		if err != nil {
			return res, err
		}
		var (
			amount int64
			price  int64
			name   string
		)
		for results.Next() {
			err = results.Scan(&amount, &price, &name)
			if err != nil {
				return res, err
			}
			res.Amount = append(res.Amount, amount)
			res.Price = append(res.Price, price)
			res.Name = append(res.Name, name)
		}
	}
	defer results.Close()
	return res, nil
}
