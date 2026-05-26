package controllers

import (
	"komik-library/config"
	"komik-library/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPeminjaman(c *gin.Context) {
	var listPeminjaman []models.Peminjaman

	config.DB.Preload("Buku").Preload("User").Find(&listPeminjaman)

	c.HTML(http.StatusOK, "index_peminjaman.html", gin.H{
		"title": "Data Peminjaman Komik",
		"data":  listPeminjaman,
	})
}
