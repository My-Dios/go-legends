package MsCustomerModel

import (
	"time"
)

type MsCustomer struct {
	CustomerID    string    `gorm:"column:customerID;type:char(36);not null;primaryKey"`
	CustomerName  string    `gorm:"column:customerName;not null"`
	DetailAddress string    `gorm:"column:detailAddress"`
	Email         string    `gorm:"column:email"`
	MobilePhone   string    `gorm:"column:mobilePhone"`
	CreatedDate   time.Time `gorm:"column:createdDate"`
	EditedDate    time.Time `gorm:"column:editedDate;autoUpdateTime"`
}

func (MsCustomer) TableName() string {
	return "ms_customer"
}
