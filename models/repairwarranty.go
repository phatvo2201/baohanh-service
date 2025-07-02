package models

import "gorm.io/gorm"

type RepairDetail struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	NgaySua          string `json:"ngaySua"`
	NoiDung          string `json:"noiDung"`
	GhiChu           string `json:"ghiChu"`
	RepairWarrantyID uint
}

type RepairWarranty struct {
	gorm.Model
	Ten     string         `json:"ten"`
	HinhAnh string         `json:"hinhAnh"`
	Sdt     string         `json:"sdt"`
	Imei    string         `json:"imei"`
	SuaChua []RepairDetail `json:"suaChua" gorm:"foreignKey:RepairWarrantyID"`
}
