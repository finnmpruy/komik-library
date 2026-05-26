package controllers

import (
	"net/http"
	"strconv"
	"time"

	"komik-library/config"
	"komik-library/models"

	"github.com/gin-gonic/gin"
)

// LIST BUKU
func IndexBuku(c *gin.Context) {
	var buku []models.Buku
	config.DB.Find(&buku)

	status := c.Query("status")
	flashMessage := ""

	if status == "create" {
		flashMessage = "Berhasil menambahkan buku baru!"
	} else if status == "update" {
		flashMessage = "Berhasil memperbarui data buku!"
	} else if status == "delete" {
		flashMessage = "Berhasil menghapus buku!"
	}

	c.HTML(200, "index.html", gin.H{
		"title": "Kelola Buku",
		"data":  buku,
		"flash": flashMessage,
	})
}

// FORM CREATE
func CreateBuku(c *gin.Context) {

	c.HTML(200, "create.html", gin.H{
		"title": "Tambah Buku",
	})
}

// STORE
func StoreBuku(c *gin.Context) {
	judul := c.PostForm("judul")
	penulis := c.PostForm("penulis")
	penerbit := c.PostForm("penerbit")
	tahunTerbit := c.PostForm("tahun_terbit")
	deskripsi := c.PostForm("deskripsi")
	stok := c.PostForm("stok")

	// Ambil file cover
	file, err := c.FormFile("cover")
	var filename string

	if err == nil {
		// Jika file ada dan tidak eror, buat nama unik dan simpan ke folder static/uploads
		filename = time.Now().Format("20060102150405") + "_" + file.Filename
		c.SaveUploadedFile(file, "static/uploads/"+filename)
	} else {
		// Jika user tidak upload gambar/eror, biarkan string kosong agar tidak bikin crash
		filename = ""
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

	// Simpan ke Database
	config.DB.Create(&buku)

	c.Redirect(http.StatusFound, "/admin/buku/?status=create")
}

// EDIT FORM
func EditBuku(c *gin.Context) {
	id := c.Param("id")

	var buku models.Buku
	config.DB.First(&buku, id)

	c.HTML(http.StatusOK, "edit.html", gin.H{
		"title": "Edit Buku",
		"data":  buku,
	})
}

// UPDATE
func UpdateBuku(c *gin.Context) {

	id := c.Param("id")

	var buku models.Buku
	config.DB.First(&buku, id)

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
		c.SaveUploadedFile(file, "static/uploads/"+filename)
		buku.Cover = filename
	}

	config.DB.Save(&buku)

	c.Redirect(http.StatusFound, "/admin/buku/?status=update")

	c.Redirect(http.StatusFound, "/admin/buku/")
}

// DELETE
func DeleteBuku(c *gin.Context) {

	id := c.Param("id")

	config.DB.Delete(&models.Buku{}, id)

	c.Redirect(http.StatusFound, "/admin/buku/?status=delete")

	c.Redirect(http.StatusFound, "/admin/buku/")
}
