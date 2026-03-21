package controllers

import (
	"encoding/json"
	"go-auth-app/config"
	"go-auth-app/models"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

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

	contentType := c.GetHeader("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		var input struct {
			Ten  string `json:"ten"`
			Sdt  string `json:"sdt"`
			Imei string `json:"imei"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if input.Ten != "" {
			rw.Ten = input.Ten
		}
		if input.Sdt != "" {
			rw.Sdt = input.Sdt
		}
		if input.Imei != "" {
			rw.Imei = input.Imei
		}
	} else {
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

		file, err := c.FormFile("hinhAnh")
		if err == nil {
			uploadPath := filepath.Join("uploads", file.Filename)
			if err := c.SaveUploadedFile(file, uploadPath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Upload image failed"})
				return
			}
			rw.HinhAnh = uploadPath
		}

		suaChuaJson := c.PostForm("suaChua")
		if suaChuaJson != "" {
			var details []models.RepairDetail
			if err := json.Unmarshal([]byte(suaChuaJson), &details); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid suaChua"})
				return
			}
			for i := range details {
				details[i].RepairWarrantyID = rw.ID
				if err := config.DB.Create(&details[i]).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Add repair detail failed"})
					return
				}
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

func AddRepairDetailToRepairWarranty(c *gin.Context) {
	id := c.Param("id")
	var rw models.RepairWarranty
	if err := config.DB.First(&rw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repair warranty not found"})
		return
	}

	var input models.RepairDetail
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = 0
	input.RepairWarrantyID = rw.ID
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Create repair detail failed"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func DeleteRepairDetailByRepairWarrantyID(c *gin.Context) {
	id := c.Param("id")
	repairID, err := strconv.ParseUint(c.Param("repairId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repair id"})
		return
	}

	result := config.DB.Where("id = ? AND repair_warranty_id = ?", uint(repairID), id).Delete(&models.RepairDetail{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Repair detail not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
