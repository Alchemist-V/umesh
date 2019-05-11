package models

import (
	u "umesh/utils"

	"github.com/jinzhu/gorm"
)

//DSRItem struct to represent DSR item.
type DSRItem struct {
	Base
	Document    string  `gorm:"unique_index:doc_item_idx"`
	ItemCode    string  `gorm:"unique_index:doc_item_idx"`
	Description string  `sql:"type:text;"`
	Unit        string  `json: "unit"`
	Amount      float64 `json: "amount"`
}

// GetItem fetches item info from DSR.
func GetItem(itemCode, dsr string) *DSRItem {
	dsrItem := &DSRItem{}

	err := GetDB().Table("dsr_items").Where("document=? and item_code = ?", dsr, itemCode).First(dsrItem).Error
	if err != nil {
		return nil
	}

	if dsrItem.ItemCode == "" {
		return nil
	}
	return dsrItem
}

// Create new DSR item
func (dsrItem *DSRItem) Create() map[string]interface{} {

	if resp, ok := dsrItem.Validate(); !ok {
		return resp
	}

	dbc := GetDB().Create(dsrItem)
	if dbc.Error != nil {
		return u.Message(false, "Exception while persisting DSR item")
	}

	response := u.Message(true, "Item registered")
	response["dsrItem"] = dsrItem
	return response
}

//Validate validates incoming create item request.
func (dsrItem *DSRItem) Validate() (map[string]interface{}, bool) {
	if len(dsrItem.Document) == 0 {
		return u.Message(false, "document not found"), false
	}

	if len(dsrItem.ItemCode) == 0 {
		return u.Message(false, "ItemId not found"), false
	}

	temp := &DSRItem{}
	err := GetDB().Table("dsr_items").Where("document = ? and item_code=?", dsrItem.Document, dsrItem.ItemCode).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.ItemCode != "" {
		return u.Message(false, "ItemId already exist for this document"), false
	}

	return u.Message(false, "Requirement passed"), true
}
