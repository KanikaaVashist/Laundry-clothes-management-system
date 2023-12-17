package models

//importing req lib

import (
	"github.com/jinzhu/gorm"
	"github.com/laundry/pkg/config"
)

var db *gorm.DB

type LaundryItem struct {
	gorm.Model
	Name     string `json:"name"`
	ItemType string `json:"item_type"`
	Status   string `json:"status"`
}

// Writing all functions

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&LaundryItem{})
}

func (li LaundryItem) CreateLaundryItem() *LaundryItem {
	db.NewRecord(li)
	db.Create(&li)
	return &li
}

func GetAllLaundryItems() []LaundryItem {
	var laundryItems []LaundryItem
	db.Find(&laundryItems)
	return laundryItems
}

func GetLaundryItemById(id int64) *LaundryItem {
	var laundryItem LaundryItem
	db.Where("ID = ?", id).First(&laundryItem)
	return &laundryItem
}

func DeleteLaundryItem(id int64) *LaundryItem {
	var laundryItem LaundryItem
	db.Where("ID = ?", id).Delete(&laundryItem)
	return &laundryItem
}

func UpdateLaundryItem(id int64, updatedItem *LaundryItem) *LaundryItem {
	var laundryItem LaundryItem
	db.Where("ID = ?", id).First(&laundryItem)

	if updatedItem.Name != "" {
		laundryItem.Name = updatedItem.Name
	}

	if updatedItem.ItemType != "" {
		laundryItem.ItemType = updatedItem.ItemType
	}

	if updatedItem.Status != "" {
		laundryItem.Status = updatedItem.Status
	}

	db.Save(&laundryItem)
	return &laundryItem
}
