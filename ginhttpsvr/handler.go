package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/allanchen1214/bingo/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var orders []Order = []Order{
	{OrderID: 1, ProductName: "phone"},
	{OrderID: 2, ProductName: "car"},
	{OrderID: 3, ProductName: "food"},
}

type Order struct {
	OrderID     uint64 `json:"orderID"`
	ProductName string `json:"productName"`
}

type OrderHandler struct {
}

func newOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func RegisterOrderRouter(r *gin.Engine) {
	h := newOrderHandler()
	g := r.Group("/order")
	g.GET("/list", h.List)
	g.GET("/item/:id", h.GetByID)
	g.GET("/item", h.GetByQueryID)
	g.POST("/add", h.Add)
}

func (h *OrderHandler) List(c *gin.Context) {
	log.Info("list all orders", zap.Any("data", orders))
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) GetByID(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	id64 := uint64(id)
	log.Info("get by id req", zap.Uint64("id", id64))
	item, err := getOrderByID(id64)
	if err != nil {
		log.Error("get by id resp", zap.Error(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Info("get by id resp", zap.Any("data", item))
	c.JSON(http.StatusOK, item)
}

func (h *OrderHandler) GetByQueryID(c *gin.Context) {
	sid := c.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(sid)
	id64 := uint64(id)
	log.Info("get by query id req", zap.Uint64("id", id64))
	item, err := getOrderByID(id64)
	if err != nil {
		log.Error("get by query id resp", zap.Error(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	log.Info("get by query id resp", zap.Any("data", item))
	c.JSON(http.StatusOK, item)
}

func getOrderByID(id uint64) (*Order, error) {
	for _, item := range orders {
		if item.OrderID == id {
			return &item, nil
		}
	}
	return nil, errors.New("NotFound")
}

func (h *OrderHandler) Add(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Error("bind json", zap.Error(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orders = append(orders, order)
	c.JSON(http.StatusOK, order)
}
