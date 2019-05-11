package models

import (
	u "umesh/utils"

	"github.com/jinzhu/gorm"
)

//CustomerBill struct to represent DSR item.
type CustomerBill struct {
	Base
	CustomerID string `json:"customer_id" gorm:"index:customer_index;type:varchar(100)"`
	Title      string `json:"title" gorm:"index:title_index;type:varchar(100)"`
	Status     string `json:"status" gorm:"type: varchar(100)"`
}

//Create create bill linked to a customer.
func (bc *CustomerBill) Create() map[string]interface{} {

	if resp, ok := bc.Validate(); !ok {
		return resp
	}

	dbc := GetDB().Create(bc)
	if dbc.Error != nil {
		return u.Message(false, "Exception while persisting bill_customer record.")
	}

	response := u.Message(true, "")
	response["billCustomer"] = bc
	return response
}

//GetAllCustomerBills fetches all bills for a customer.
func GetAllCustomerBills(cid string) []CustomerBill {
	cb := []CustomerBill{}
	GetDB().Table("customer_bills").Where("customer_id=?", cid).Find(&cb)
	return cb
}

//GetBillByTitle fetches bill by title for a customer.
func GetBillByTitle(cid string, t string) *CustomerBill {
	cb := &CustomerBill{}
	GetDB().Table("customer_bills").Where("customer_id=? and title=?", cid, t).First(cb)
	return cb
}

//Validate validates instance of customer bill.
func (bc *CustomerBill) Validate() (map[string]interface{}, bool) {
	if len(bc.CustomerID) == 0 {
		return u.Message(false, "CustomerID not found"), false
	}
	temp := &CustomerBill{}
	err := GetDB().Table("customer_bills").Where("customer_id=? and title = ?", bc.CustomerID, bc.Title).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Title already in use for billId: "+temp.ID.String()), false
	}

	if len(temp.Title) != 0 {
		return u.Message(false, "Title already in use for billId: "+temp.ID.String()), false
	}

	return u.Message(false, "Requirement passed"), true
}
