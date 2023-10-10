package MsCustomerMigration

import (
	MsCustomerModel "github.com/My-Dios/go-legends/models/ms-customer"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) error {
	return db.AutoMigrate(&MsCustomerModel.MsCustomer{})
}

func Drop(db *gorm.DB) error {
	return db.Migrator().DropTable(&MsCustomerModel.MsCustomer{})
}
