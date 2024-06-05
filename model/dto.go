package model

type OrderRequest struct {
	NoTable       int64         `json:"noTable"`
	MenuMakanan   []MenuMakanan `json:"menuMakanan"`
}

type OrderUpdateRequest struct {
	Id           string       `json:"id"`
	NoTable      int64        `json:"noTable"`
	MenuMakanan  []MenuMakanan `json:"menuMakanan"`
}

type MenuMakanan struct {
	IdMakanan int64 `json:"idMakanan"`
	Qty       int64 `json:"qty"` 
}