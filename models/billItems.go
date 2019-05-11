package models

import (
	u "umesh/utils"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//BillItem struct to represent a bill item.
type BillItem struct {
	Base
	BillID                  uuid.UUID `json:"bill_id" gorm:"index:customer_bill_index;type:varchar(36)"`
	ItemID                  uuid.UUID `json:"item_id" gorm:"index:item_id_index;type:varchar(36)"`
	Unit                    string    `json:"unit"`
	Quantity                float64   `json:"quantity"`
	Rate                    float64   `json:"rate"`
	AdjustmentFactorPercent float64   `json:"adjustment_factor_percent"`
	AdjustmentFactorSign    int       `json:"adjustment_factor_sign"`
	Comment                 string    `json:"comment"`
}

// Create registes an item for a bill
func (bi *BillItem) Create() map[string]interface{} {
	if resp, ok := bi.Validate(); !ok {
		return resp
	}

	dbc := GetDB().Create(bi)
	if dbc.Error != nil {
		return u.Message(false, "Exception while persisting DSR item")
	}

	response := u.Message(true, "Item registered")
	response["billItem"] = bi
	return response
}

//GetBillItems fetches all bill items for a bill.
func GetBillItems(bId uuid.UUID) []BillItem {
	items := []BillItem{}
	db := GetDB().Table("bill_items").Where("bill_id=?", bId).Find(&items)

	if db.Error != nil {
		return nil
	}
	return items
}

//Validate validates item before linking it with a bill
func (bi *BillItem) Validate() (map[string]interface{}, bool) {

	if len(bi.BillID.String()) == 0 {
		return u.Message(false, "BillID not found"), false
	}

	if len(bi.ItemID.String()) == 0 {
		return u.Message(false, "ItemID not found"), false
	}

	if bi.Quantity == 0.0 {
		return u.Message(false, "Quantity not found"), false
	}

	temp := &BillItem{}
	err := GetDB().Table("bill_items").Where("bill_id = ? and item_id=?", bi.BillID, bi.ItemID).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	return u.Message(false, "Requirement passed"), true
}
