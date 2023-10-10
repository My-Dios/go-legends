package MsItemController

import (
	"math"
	"net/http"
	"strconv"

	ResponseAPIHelper "github.com/My-Dios/go-legends/Helpers/response-api"
	"github.com/My-Dios/go-legends/config/db"
	MsItemModel "github.com/My-Dios/go-legends/models/ms-item"
	MsItemModelForm "github.com/My-Dios/go-legends/modules/v1/models/ms-item"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db.InitDB()
	items, err := MsItemModelForm.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	count, err := MsItemModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	limitParam := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'limit' parameter"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	totalPages := int(math.Ceil(float64(count) / float64(limit)))
	paginatedItems := MsItemModelForm.PaginateItems(items, page, limit)
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Get All Item Successfully", c, paginatedItems, int(count), limit, page, totalPages)
}

func Show(c *gin.Context) {
	db.InitDB()
	itemID := c.Param("id")
	count, err := MsItemModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	item, err := MsItemModelForm.GetData(itemID)
	if err != nil {
		ResponseAPIHelper.GetStandarization(http.StatusNotFound, "Item Not Found", c, []MsItemModel.MsItem{}, int(count), 1, 1, 1)
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Get Item Successfully", c, item, int(count), 1, 1, 1)
}

func Create(c *gin.Context) {
	db.InitDB()
	var newItem MsItemModel.MsItem
	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MsItemModelForm.CreateData(&newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	count, err := MsItemModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusCreated, "Create Item Successfully", c, newItem, int(count), 1, 1, 1)
}

func Update(c *gin.Context) {
	db.InitDB()
	itemID := c.Param("id")
	var updatedItem MsItemModel.MsItem
	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MsItemModelForm.UpdateData(itemID, &updatedItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	count, err := MsItemModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Update Item Successfully", c, updatedItem, int(count), 1, 1, 1)
}

func Delete(c *gin.Context) {
	db.InitDB()
	itemID := c.Param("id")
	count, err := MsItemModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	deletedItem, err := MsItemModelForm.GetData(itemID)
	if err != nil {
		ResponseAPIHelper.GetStandarization(http.StatusNotFound, "Item Not Found", c, []MsItemModel.MsItem{}, int(count), 1, 1, 1)
		return
	}
	if err := MsItemModelForm.DeleteData(itemID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Delete Item Successfully", c, deletedItem, int(count), 1, 1, 1)
}
