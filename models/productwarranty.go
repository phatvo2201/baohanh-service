package models

import "gorm.io/gorm"

type RepairData struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	RepairDate        string `json:"repairDate"`
	WarrantyEnd       string `json:"warrantyEnd"`
	Content           string `json:"content"`
	Note              string `json:"note"`
	ProductWarrantyID uint
}

type ProductWarranty struct {
	gorm.Model
	Ten        string       `json:"ten"`
	HinhAnh    string       `json:"hinhAnh"`
	Imei       string       `json:"imei"`
	NgayMua    string       `json:"ngayMua"`
	HetHan     string       `json:"hetHan"`
	RepairData []RepairData `json:"repairData" gorm:"foreignKey:ProductWarrantyID"`
}
