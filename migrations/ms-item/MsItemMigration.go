package MsItemMigration

import (
	MsItemModel "github.com/My-Dios/go-legends/models/ms-item"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) error {
	return db.AutoMigrate(&MsItemModel.MsItem{})
}

func Drop(db *gorm.DB) error {
	return db.Migrator().DropTable(&MsItemModel.MsItem{})
}
