package model

import "time"

type Orders struct {
	ID 			string      `json:"_id" bson:"_id,omitempty"`
	NoTable 	int64       `json:"noTable" bson:"noTable"`
	MenuOrder   []MenuOrder `json:"menuMakanan" bson:"menuMakanan"`
	CreatedAt 	time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt 	time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type OrderUpdate struct {
	NoTable 	int64       `json:"noTable" bson:"noTable"`
	MenuOrder   []MenuOrder `json:"menuMakanan" bson:"menuMakanan"`
	UpdatedAt 	time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type MenuOrder struct {
	IdMakanan int64 `json:"idMakanan" bson:"idMakanan"`
	Qty       int64 `json:"qty" bson:"qty"` 
}

type OrderResponse struct {
	Code            int64         `json:"code"`
	OrderResult     []OrderResult `json:"result"`
	Message         string 	      `json:"message"`
}

type OrderResult struct {
	NoTable 	       int64   	         `json:"noTable" bson:"noTable"`
	OrderMenuResponse  []OrderMenuResponse `json:"orderMenu"`
	TotalPrice  	   int64 			 `json:"totalPrice"`
	DateOrder   	   time.Time         `json:"dateOrder"` 
}

type OrderMenuResponse struct {
	IdMakanan int64 `json:"idMakanan" bson:"idMakanan"`
	Qty       int64 `json:"qty" bson:"qty"` 
	Price     int64 `json:"price"`
}

func NewOrderMapping(payload OrderRequest) Orders  {
	menuOrder := MenuOrder{}
	menusOrder := []MenuOrder{}
	for _, v := range payload.MenuMakanan {
		menuOrder.IdMakanan = v.IdMakanan
		menuOrder.Qty = v.Qty
		menusOrder = append(menusOrder, menuOrder)
	}

	return Orders{
		NoTable: payload.NoTable,
		MenuOrder: menusOrder,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewMappingOrderResponse(payload OrderRequest) OrderResult {
	menuOrder := OrderMenuResponse{}
	menusOrder := []OrderMenuResponse{}
	for _, v := range payload.MenuMakanan {
		menuOrder.IdMakanan = v.IdMakanan
		menuOrder.Qty = v.Qty
		menusOrder = append(menusOrder, menuOrder)
	}
	return OrderResult{
		NoTable: payload.NoTable,
		OrderMenuResponse: menusOrder,
		DateOrder: time.Now(),
	}
}