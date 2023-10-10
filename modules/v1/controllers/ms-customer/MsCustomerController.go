package MsCustomerController

import (
	"math"
	"net/http"
	"strconv"

	ResponseAPIHelper "github.com/My-Dios/go-legends/Helpers/response-api"
	"github.com/My-Dios/go-legends/config/db"
	MsCustomerModel "github.com/My-Dios/go-legends/models/ms-customer"
	MsCustomerModelForm "github.com/My-Dios/go-legends/modules/v1/models/ms-customer"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	db.InitDB()
	customers, err := MsCustomerModelForm.GetAllData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	count, err := MsCustomerModelForm.GetCountData()
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
	paginatedItems := MsCustomerModelForm.PaginateItems(customers, page, limit)
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Get All Customer Successfully", c, paginatedItems, int(count), limit, page, totalPages)
}

func Show(c *gin.Context) {
	db.InitDB()
	customerID := c.Param("id")
	customer, err := MsCustomerModelForm.GetData(customerID)
	if err != nil {
		ResponseAPIHelper.GetStandarization(http.StatusNotFound, "Customer Not Found", c, []MsCustomerModel.MsCustomer{}, 0, 1, 1, 1)
		return
	}
	count, err := MsCustomerModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Get Customer Successfully", c, customer, int(count), 1, 1, 1)
}

func Create(c *gin.Context) {
	db.InitDB()
	var newCustomer MsCustomerModel.MsCustomer
	if err := c.BindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MsCustomerModelForm.CreateData(&newCustomer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	count, err := MsCustomerModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusCreated, "Create Customer Successfully", c, newCustomer, int(count), 1, 1, 1)
}

func Update(c *gin.Context) {
	db.InitDB()
	customerID := c.Param("id")
	var updatedCustomer MsCustomerModel.MsCustomer
	if err := c.BindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := MsCustomerModelForm.UpdateData(customerID, &updatedCustomer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	count, err := MsCustomerModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Update Customer Successfully", c, updatedCustomer, int(count), 1, 1, 1)
}

func Delete(c *gin.Context) {
	db.InitDB()
	customerID := c.Param("id")
	count, err := MsCustomerModelForm.GetCountData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	deletedCustomer, err := MsCustomerModelForm.GetData(customerID)
	if err != nil {
		ResponseAPIHelper.GetStandarization(http.StatusNotFound, "Customer Not Found", c, []MsCustomerModel.MsCustomer{}, int(count), 1, 1, 1)
		return
	}
	if err := MsCustomerModelForm.DeleteData(customerID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ResponseAPIHelper.GetStandarization(http.StatusOK, "Delete Customer Successfully", c, deletedCustomer, int(count), 1, 1, 1)
}
