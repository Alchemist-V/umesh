package controllers

import (
	"encoding/json"
	"net/http"
	"umesh/models"
	u "umesh/utils"

	uuid "github.com/satori/go.uuid"
)

// CreateCustomerBill creates a customer bill
var CreateCustomerBill = func(w http.ResponseWriter, r *http.Request) {
	cb := &models.CustomerBill{}
	err := json.NewDecoder(r.Body).Decode(cb)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := cb.Create()
	u.Respond(w, resp)
}

//RegisterItemToBill adds a row to bill with given item.
var RegisterItemToBill = func(w http.ResponseWriter, r *http.Request) {
	bi := &models.BillItem{}
	err := json.NewDecoder(r.Body).Decode(bi)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := bi.Create()
	u.Respond(w, resp)
}

//GetAllBillsByCustomer fetch all bills by customer.
var GetAllBillsByCustomer = func(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	cid := keys.Get("cid")
	cbs := models.GetAllCustomerBills(cid)
	resp := u.Message(true, "success")
	resp["bills"] = cbs
	u.Respond(w, resp)
}

//GetBillByTitle fetches bill by customer and bill title.
var GetBillByTitle = func(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	cid := keys.Get("cid")
	t := keys.Get("title")
	cb := models.GetBillByTitle(cid, t)
	resp := u.Message(true, "success")
	resp["bill"] = cb
	u.Respond(w, resp)
}

//GetBillItems fetch all items in a particular bill.
var GetBillItems = func(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	uid, err := uuid.FromString(keys.Get("bid"))
	if err != nil {
		resp := u.Message(false, "Failure extracting bill id from query param, please ensure you are passing a valid uuid.")
		u.Respond(w, resp)
	}
	items := models.GetBillItems(uid)
	resp := u.Message(true, "success")
	resp["item"] = items
	u.Respond(w, resp)
}
