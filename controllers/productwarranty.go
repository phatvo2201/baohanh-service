package controllers

import (
	"encoding/json"
	"go-auth-app/config"
	"go-auth-app/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CreateProductWarranty(c *gin.Context) {
	var pw models.ProductWarranty

	// Nhận các trường text
	pw.Ten = c.PostForm("ten")
	pw.Imei = c.PostForm("imei")
	pw.NgayMua = c.PostForm("ngayMua")
	pw.HetHan = c.PostForm("hetHan")

	// Nhận file hình ảnh
	file, err := c.FormFile("hinhAnh")
	if err == nil {
		uploadPath := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload image failed"})
			return
		}
		pw.HinhAnh = uploadPath
	}

	// Nhận trường repairData (JSON string)
	repairDataJson := c.PostForm("repairData")
	if repairDataJson != "" {
		var details []models.RepairData
		if err := json.Unmarshal([]byte(repairDataJson), &details); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repairData"})
			return
		}
		pw.RepairData = details
	}

	if err := config.DB.Create(&pw).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, pw)
}

func GetAllProductWarranties(c *gin.Context) {
	var pws []models.ProductWarranty
	if err := config.DB.Preload("RepairData").Find(&pws).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fetch failed"})
		return
	}
	c.JSON(http.StatusOK, pws)
}

func GetProductWarrantyByID(c *gin.Context) {
	id := c.Param("id")
	var pw models.ProductWarranty
	if err := config.DB.Preload("RepairData").First(&pw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, pw)
}

func UpdateProductWarranty(c *gin.Context) {
	id := c.Param("id")
	var pw models.ProductWarranty
	if err := config.DB.Preload("RepairData").First(&pw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	// Nhận các trường text
	ten := c.PostForm("ten")
	imei := c.PostForm("imei")
	ngayMua := c.PostForm("ngayMua")
	hetHan := c.PostForm("hetHan")
	if ten != "" {
		pw.Ten = ten
	}
	if imei != "" {
		pw.Imei = imei
	}
	if ngayMua != "" {
		pw.NgayMua = ngayMua
	}
	if hetHan != "" {
		pw.HetHan = hetHan
	}

	// Nhận file hình ảnh mới (nếu có)
	file, err := c.FormFile("hinhAnh")
	if err == nil {
		uploadPath := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload image failed"})
			return
		}
		pw.HinhAnh = uploadPath
	}

	// Nhận trường repairData mới (JSON string) và thêm vào
	repairDataJson := c.PostForm("repairData")
	if repairDataJson != "" {
		var details []models.RepairData
		if err := json.Unmarshal([]byte(repairDataJson), &details); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repairData"})
			return
		}
		// Gán ProductWarrantyID cho từng bản ghi mới và thêm vào DB
		for i := range details {
			details[i].ProductWarrantyID = pw.ID
			if err := config.DB.Create(&details[i]).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Add repair data failed"})
				return
			}
		}
	}

	if err := config.DB.Save(&pw).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, pw)
}

func DeleteProductWarranty(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.ProductWarranty{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func ListRepairDetailsByProductWarrantyID(c *gin.Context) {
	id := c.Param("id")
	var details []models.RepairData
	if err := config.DB.Where("product_warranty_id = ?", id).Find(&details).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fetch failed"})
		return
	}
	c.JSON(http.StatusOK, details)
}
