package models

type UriId struct {
	ID int `uri:"id" binding:"required"`
}

type CustomerParam struct {
	ID           uint    `json:"id"`
	Nik          string  `json:"nik"`
	FullName     string  `json:"full_name"`
	LegalName    string  `json:"legal_name"`
	TempatLahir  string  `json:"tempat_lahir"`
	TanggalLahir string  `json:"tanggal_lahir"`
	Gaji         float64 `json:"gaji"`
	FotoKtp      string  `json:"foto_ktp"`
	FotoSelfi    string  `json:"foto_selfi"`
}

type ComponentServerSide struct {
	Limit     int    `json:"limit"`
	Skip      int    `json:"skip"`
	SortType  string `json:"sort_type"`
	SortBy    string `json:"sort_by"`
	Search    string `json:"search"`
	Offset    int    `json:"offset"`
	Condition string `json:"condition"`
	From      string `json:"from"`
	To        string `json:"to"`
}
