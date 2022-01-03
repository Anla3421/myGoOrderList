package controller

import (
	"fmt"
	"server/api/protocol"
	"server/env"
	"server/model/dao/shoppingCartQuery"
	"server/service/pbclient"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	ShoppingCartQueryReq struct {
		Jwt string `form:"jwt" json:"jwt" binding:"required"`
	}

	ShoppingCartQueryStorage struct {
		Err     error
		BuyerID int
		Account string
	}

	ShoppingCartQueryTask struct {
		Req     *ShoppingCartQueryReq
		Res     *protocol.Response
		Storage *ShoppingCartQueryStorage
	}
)

// NewShopingCartTask:實體化Task
func NewShoppingCartQueryTask() *ShoppingCartQueryTask {
	return &ShoppingCartQueryTask{
		Req:     &ShoppingCartQueryReq{},
		Res:     &protocol.Response{},
		Storage: &ShoppingCartQueryStorage{},
	}
}

// shopingCart
func ShoppingCartQuery(c *gin.Context) {
	task := NewShoppingCartQueryTask()
	c.Set(env.APIResKeyInGinContext, task.Res)

	if shouldBreak := task.ShouldBind(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.AccountJwtCheck(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	if shouldBreak := task.ShoppingCartQuery(c); shouldBreak {
		c.Error(task.Storage.Err)
		return
	}
	return
}

// ShouldBind:參數解析
// 檢查input的資料key的型態與名稱正確性，Value不可為空
func (task *ShoppingCartQueryTask) ShouldBind(c *gin.Context) bool {
	if err := c.ShouldBindBodyWith(task.Req, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		task.Storage.Err = err
		return true
	}
	return false
}

// AccountJwtCheck:
func (task *ShoppingCartQueryTask) AccountJwtCheck(c *gin.Context) bool {
	res := pbclient.GetMemberIDByJWT(task.Req.Jwt)
	if res.ID == 0 {
		task.Storage.Err = fmt.Errorf("Please login first")
		return true
	}
	task.Storage.BuyerID = int(res.ID)
	task.Storage.Account = res.Account
	return false
}

func (task *ShoppingCartQueryTask) ShoppingCartQuery(c *gin.Context) bool {
	task.Res.Result, task.Storage.Err = shoppingCartQuery.ShoppingCartQuery(task.Storage.BuyerID, task.Storage.Account)
	if task.Storage.Err != nil {
		return true
	}
	return false
}
