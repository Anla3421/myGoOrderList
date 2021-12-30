package controller

import (
	"database/sql"
	"fmt"
	"server/api/protocol"
	"server/env"
	"server/model/dao/shoppingCartOrder"
	"server/model/dto"
	"server/service/pbclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	ShoppingCartReq struct {
		Jwt        string `form:"jwt" json:"jwt" binding:"required"`
		TotalPrice int    `form:"totalPrice" json:"totalPrice" binding:"required"`

		ItemID []int `form:"itemID" json:"itemID" binding:"required"`
		Amount []int `form:"amount" json:"amount" binding:"required"`
		Price  []int `form:"price" json:"price" binding:"required"`
		ID     int   `form:"id" json:"id"`
	}

	ShoppingCartStorage struct {
		Err     error
		BuyerID int
		Account string
	}

	ShoppingCartTask struct {
		Req     *ShoppingCartReq
		Res     *protocol.Response
		Storage *ShoppingCartStorage
	}
)

var mysqlConn *sql.DB

func init() {
	CreateConn()
}

func CreateConn() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}
	mysqlConn = db
}

// NewShopingCartTask:實體化Task
func NewShopingCartTask() *ShoppingCartTask {
	return &ShoppingCartTask{
		Req:     &ShoppingCartReq{},
		Res:     &protocol.Response{},
		Storage: &ShoppingCartStorage{},
	}
}

// shopingCart
func ShopingCartOrder(c *gin.Context) {
	task := NewShopingCartTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.AccountJwtCheck(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.ShoppingCartOrder(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	return
}

// ShouldBind:參數解析
// 檢查input的資料key的型態與名稱正確性，Value不可為空
func (task *ShoppingCartTask) ShouldBind(c *gin.Context) bool {
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

// AccountJwtCheck:
func (task *ShoppingCartTask) AccountJwtCheck(c *gin.Context) bool {
	res := pbclient.GetMemberIDByJWT(task.Req.Jwt)
	if res.ID == 0 {
		task.Storage.Err = fmt.Errorf("Please login first")
		return true
	}
	task.Storage.BuyerID = int(res.ID)
	task.Storage.Account = res.Account
	return false
}

func (task *ShoppingCartTask) ShoppingCartOrder(c *gin.Context) bool {
	t := time.Now()
	timeNow := t.Format("2006-01-02 15:04:05")
	fmt.Println(timeNow)
	task.Storage.Err = shoppingCartOrder.ShoppingCartOrder(&dto.ShoppingCartOrder{
		TotalPrice: task.Req.TotalPrice,
		ItemID:     task.Req.ItemID,
		Amount:     task.Req.Amount,
		Price:      task.Req.Price,
		BuyerID:    task.Storage.BuyerID,
		TimeNow:    timeNow,
	})
	if task.Storage.Err != nil {
		return true
	}
	return false
}
