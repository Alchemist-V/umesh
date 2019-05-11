package controllers

import (
	"encoding/json"
	"net/http"
	"umesh/models"
	u "umesh/utils"
)

// FetchItem fetches item recordf
var FetchItem = func(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	source := keys.Get("src")
	item := keys.Get("item")

	data := models.GetItem(item, source)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// CreateItem creates new item record.
var CreateItem = func(w http.ResponseWriter, r *http.Request) {
	dsrItem := &models.DSRItem{}
	err := json.NewDecoder(r.Body).Decode(dsrItem)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := dsrItem.Create()
	u.Respond(w, resp)
}
