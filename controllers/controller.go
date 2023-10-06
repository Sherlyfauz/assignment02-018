package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "Home")
}

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newOrder,
	})

}

func GetOrders(ctx *gin.Context) {
	var getOrders []models.Order

	if err := database.DB.Preload("Items").Find(&getOrders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": getOrders,
	})
}

func GetOrderById(ctx *gin.Context) {
	var orderItem models.Order
	orderIDStr := ctx.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Preload("Items").First(&orderItem, "id = ?", orderID).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orderItem,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("id")
	orderID, _ := strconv.Atoi(orderIDStr)

	var order models.Order
	if err := database.DB.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := range order.Items {
		itemID := order.Items[i].ID // ID atau primary key item

		if err := database.DB.Model(&models.Item{}).Where("id = ?", itemID).Updates(&order.Items[i]).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Untuk Update Item!", "details": err.Error()})
			return
		}
	}

	if err := database.DB.Save(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Untuk Update Item!", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	orderIDStr := ctx.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order Anda Telah Berhasil Dihapus!"})
}
