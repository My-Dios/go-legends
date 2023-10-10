package MsItemModel

import (
	"time"
)

type MsItem struct {
	ItemID      string    `gorm:"column:itemID;type:char(36);not null;primaryKey"`
	ItemName    string    `gorm:"column:itemName;not null"`
	Price       float64   `gorm:"column:price;type:decimal(14,4);default:0.0000"`
	CreatedDate time.Time `gorm:"column:createdDate"`
	EditedDate  time.Time `gorm:"column:editedDate;autoUpdateTime"`
}

func (MsItem) TableName() string {
	return "ms_item"
}
