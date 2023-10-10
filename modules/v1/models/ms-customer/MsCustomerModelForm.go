package MsCustomerModelForm

import (
	"time"

	"github.com/My-Dios/go-legends/config/db"
	MsCustomerModel "github.com/My-Dios/go-legends/models/ms-customer"
	"github.com/google/uuid"
)

func generateUUID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuidObj.String(), nil
}

func GetAllData() ([]MsCustomerModel.MsCustomer, error) {
	var customers []MsCustomerModel.MsCustomer
	if err := db.GetDB().Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func GetData(customerID string) (MsCustomerModel.MsCustomer, error) {
	var customer MsCustomerModel.MsCustomer
	if err := db.GetDB().Where("customerID = ?", customerID).First(&customer).Error; err != nil {
		return MsCustomerModel.MsCustomer{}, err
	}
	return customer, nil
}

func CreateData(newCustomer *MsCustomerModel.MsCustomer) error {
	newID, err := generateUUID()
	if err != nil {
		return err
	}
	newCustomer.CustomerID = newID
	newCustomer.CreatedDate = time.Now()
	newCustomer.EditedDate = time.Now()
	if err := db.GetDB().Create(newCustomer).Error; err != nil {
		return err
	}
	return nil
}

func GetCountData() (int64, error) {
	var count int64
	if err := db.GetDB().Model(&MsCustomerModel.MsCustomer{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func UpdateData(customerID string, updatedCustomer *MsCustomerModel.MsCustomer) error {
	var existingCustomer MsCustomerModel.MsCustomer
	if err := db.GetDB().Where("customerID = ?", customerID).First(&existingCustomer).Error; err != nil {
		return err
	}
	existingCustomer.CustomerName = updatedCustomer.CustomerName
	existingCustomer.DetailAddress = updatedCustomer.DetailAddress
	existingCustomer.Email = updatedCustomer.Email
	existingCustomer.MobilePhone = updatedCustomer.MobilePhone
	existingCustomer.EditedDate = time.Now()
	if err := db.GetDB().Save(&existingCustomer).Error; err != nil {
		return err
	}
	return nil
}

func DeleteData(customerID string) error {
	var existingCustomer MsCustomerModel.MsCustomer
	if err := db.GetDB().Where("customerID = ?", customerID).First(&existingCustomer).Error; err != nil {
		return err
	}
	if err := db.GetDB().Delete(&existingCustomer).Error; err != nil {
		return err
	}
	return nil
}

func PaginateItems(customers []MsCustomerModel.MsCustomer, currentPage, limit int) []MsCustomerModel.MsCustomer {
	startIndex := (currentPage - 1) * limit
	endIndex := startIndex + limit
	if startIndex >= len(customers) {
		return []MsCustomerModel.MsCustomer{}
	}
	if endIndex > len(customers) {
		endIndex = len(customers)
	}
	return customers[startIndex:endIndex]
}
