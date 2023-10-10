package MsItemModelForm

import (
	"fmt"
	"time"

	"github.com/My-Dios/go-legends/config/db"
	MsItemModel "github.com/My-Dios/go-legends/models/ms-item"
)

func GetAllData() ([]MsItemModel.MsItem, error) {
	var items []MsItemModel.MsItem
	if err := db.GetDB().Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func GetData(itemID string) (MsItemModel.MsItem, error) {
	var item MsItemModel.MsItem
	if err := db.GetDB().Where("itemID = ?", itemID).First(&item).Error; err != nil {
		return MsItemModel.MsItem{}, err
	}
	return item, nil
}

func GetCountData() (int64, error) {
	var count int64
	if err := db.GetDB().Model(&MsItemModel.MsItem{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func getIncrementedID() (string, error) {
	var maxID string
	if err := db.GetDB().Model(&MsItemModel.MsItem{}).Select("itemID").
		Order("CAST(SUBSTRING_INDEX(itemID, '-', -1) AS SIGNED) DESC").Limit(1).
		Pluck("itemID", &maxID).Error; err != nil {
		return "", err
	}
	maxIDNumber := 0
	if maxID != "" {
		_, err := fmt.Sscanf(maxID, "ITM-%03d", &maxIDNumber)
		if err != nil {
			return "", err
		}
	}
	newIDNumber := maxIDNumber + 1
	newID := fmt.Sprintf("ITM-%03d", newIDNumber)
	return newID, nil
}

func CreateData(newItem *MsItemModel.MsItem) error {
	newID, err := getIncrementedID()
	if err != nil {
		return err
	}
	newItem.ItemID = newID
	newItem.CreatedDate = time.Now()
	newItem.EditedDate = time.Now()
	if err := db.GetDB().Create(newItem).Error; err != nil {
		return err
	}
	return nil
}

func UpdateData(itemID string, updatedItem *MsItemModel.MsItem) error {
	var existingItem MsItemModel.MsItem
	if err := db.GetDB().Where("itemID = ?", itemID).First(&existingItem).Error; err != nil {
		return err
	}
	existingItem.ItemName = updatedItem.ItemName
	existingItem.Price = updatedItem.Price
	existingItem.EditedDate = time.Now()
	if err := db.GetDB().Save(&existingItem).Error; err != nil {
		return err
	}
	return nil
}

func DeleteData(itemID string) error {
	var existingItem MsItemModel.MsItem
	if err := db.GetDB().Where("itemID = ?", itemID).First(&existingItem).Error; err != nil {
		return err
	}
	if err := db.GetDB().Delete(&existingItem).Error; err != nil {
		return err
	}
	return nil
}

func PaginateItems(items []MsItemModel.MsItem, currentPage, limit int) []MsItemModel.MsItem {
	startIndex := (currentPage - 1) * limit
	endIndex := startIndex + limit
	if startIndex >= len(items) {
		return []MsItemModel.MsItem{}
	}
	if endIndex > len(items) {
		endIndex = len(items)
	}
	return items[startIndex:endIndex]
}
