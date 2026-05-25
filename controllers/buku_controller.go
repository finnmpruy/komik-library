package controllers

import (
	"net/http"
	"strconv"
	"time"

	"komik-library/models"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Komik Library",
	})
}

func IndexBuku(c *gin.Context) {
	var buku []models.Buku
	db.Find(&buku)

	c.HTML(http.StatusOK, "buku/index.html", gin.H{
		"title": "Data Buku",
		"data":  buku,
	})
}

func CreateBuku(c *gin.Context) {
	c.HTML(http.StatusOK, "buku/create.html", gin.H{
		"title": "Tambah Buku",
	})
}

func StoreBuku(c *gin.Context) {

	judul := c.PostForm("judul")
	penulis := c.PostForm("penulis")
	penerbit := c.PostForm("penerbit")
	tahunTerbit := c.PostForm("tahun_terbit")
	deskripsi := c.PostForm("deskripsi")
	stok := c.PostForm("stok")

	file, err := c.FormFile("cover")
	var filename string

	if err == nil {
		filename = time.Now().Format("20060102150405") + "_" + file.Filename
		path := "uploads/" + filename
		c.SaveUploadedFile(file, path)
	}

	stokInt, _ := strconv.Atoi(stok)

	buku := models.Buku{
		Judul:       judul,
		Penulis:     penulis,
		Penerbit:    penerbit,
		TahunTerbit: tahunTerbit,
		Deskripsi:   deskripsi,
		Stok:        stokInt,
		Cover:       filename,
	}

	db.Create(&buku)

	c.Redirect(http.StatusFound, "/buku")
}

func ShowBuku(c *gin.Context) {

	id := c.Param("id")

	var buku models.Buku
	db.First(&buku, id)

	c.HTML(http.StatusOK, "buku/show.html", gin.H{
		"title": "Detail Buku",
		"data":  buku,
	})
}

func EditBuku(c *gin.Context) {

	id := c.Param("id")

	var buku models.Buku
	db.First(&buku, id)

	c.HTML(http.StatusOK, "buku/edit.html", gin.H{
		"title": "Edit Buku",
		"data":  buku,
	})
}

func UpdateBuku(c *gin.Context) {

	id := c.Param("id")

	var buku models.Buku
	db.First(&buku, id)

	buku.Judul = c.PostForm("judul")
	buku.Penulis = c.PostForm("penulis")
	buku.Penerbit = c.PostForm("penerbit")
	buku.TahunTerbit = c.PostForm("tahun_terbit")
	buku.Deskripsi = c.PostForm("deskripsi")

	stok := c.PostForm("stok")
	stokInt, _ := strconv.Atoi(stok)
	buku.Stok = stokInt

	file, err := c.FormFile("cover")
	if err == nil {
		filename := time.Now().Format("20060102150405") + "_" + file.Filename
		path := "uploads/" + filename

		c.SaveUploadedFile(file, path)
		buku.Cover = filename
	}

	db.Save(&buku)

	c.Redirect(http.StatusFound, "/buku")
}

func DeleteBuku(c *gin.Context) {

	id := c.Param("id")

	db.Delete(&models.Buku{}, id)

	c.Redirect(http.StatusFound, "/buku")
}
