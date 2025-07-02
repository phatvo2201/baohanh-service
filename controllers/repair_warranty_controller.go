package controllers

import (
	"encoding/json"
	"go-auth-app/config"
	"go-auth-app/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CreateRepairWarranty(c *gin.Context) {
	var rw models.RepairWarranty

	// Nhận các trường text
	rw.Ten = c.PostForm("ten")
	rw.Sdt = c.PostForm("sdt")
	rw.Imei = c.PostForm("imei")

	// Nhận file hình ảnh
	file, err := c.FormFile("hinhAnh")
	if err == nil {
		uploadPath := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload image failed"})
			return
		}
		rw.HinhAnh = uploadPath
	}

	// Nhận trường sửa chữa (JSON string)
	suaChuaJson := c.PostForm("suaChua")
	if suaChuaJson != "" {
		var details []models.RepairDetail
		if err := json.Unmarshal([]byte(suaChuaJson), &details); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid suaChua"})
			return
		}
		rw.SuaChua = details
	}

	if err := config.DB.Create(&rw).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create failed"})
		return
	}
	c.JSON(http.StatusCreated, rw)
}

func GetAllRepairWarranties(c *gin.Context) {
	var rws []models.RepairWarranty
	if err := config.DB.Preload("SuaChua").Find(&rws).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fetch failed"})
		return
	}
	c.JSON(http.StatusOK, rws)
}

func GetRepairWarrantyByID(c *gin.Context) {
	id := c.Param("id")
	var rw models.RepairWarranty
	if err := config.DB.Preload("SuaChua").First(&rw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	c.JSON(http.StatusOK, rw)
}

func UpdateRepairWarranty(c *gin.Context) {
	id := c.Param("id")
	var rw models.RepairWarranty
	if err := config.DB.Preload("SuaChua").First(&rw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	// Nhận các trường text
	ten := c.PostForm("ten")
	sdt := c.PostForm("sdt")
	imei := c.PostForm("imei")
	if ten != "" {
		rw.Ten = ten
	}
	if sdt != "" {
		rw.Sdt = sdt
	}
	if imei != "" {
		rw.Imei = imei
	}

	// Nhận file hình ảnh mới (nếu có)
	file, err := c.FormFile("hinhAnh")
	if err == nil {
		uploadPath := filepath.Join("uploads", file.Filename)
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload image failed"})
			return
		}
		rw.HinhAnh = uploadPath
	}

	// Nhận trường sửa chữa mới (JSON string) và thêm vào
	suaChuaJson := c.PostForm("suaChua")
	if suaChuaJson != "" {
		var details []models.RepairDetail
		if err := json.Unmarshal([]byte(suaChuaJson), &details); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid suaChua"})
			return
		}
		// Gán RepairWarrantyID cho từng bản ghi mới và thêm vào DB
		for i := range details {
			details[i].RepairWarrantyID = rw.ID
			if err := config.DB.Create(&details[i]).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Add repair detail failed"})
				return
			}
		}
	}

	if err := config.DB.Save(&rw).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(http.StatusOK, rw)
}

func DeleteRepairWarranty(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.RepairWarranty{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func ListRepairDetailsByRepairWarrantyID(c *gin.Context) {
	id := c.Param("id")
	var details []models.RepairDetail
	if err := config.DB.Where("repair_warranty_id = ?", id).Find(&details).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fetch failed"})
		return
	}
	c.JSON(http.StatusOK, details)
}
