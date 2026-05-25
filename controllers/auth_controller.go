package controllers

import (
	"net/http"

	"komik-library/config"
	"komik-library/models"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})

}
func Login(c *gin.Context) {

	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User

	config.DB.Where("email = ? AND password = ?", email, password).First(&user)

	// JIKA USER TIDAK DITEMUKAN
	if user.ID == 0 {

		c.String(http.StatusUnauthorized, "Email atau password salah")

		return
	}

	// JIKA ROLE ADMIN
	if user.Role == "admin" {

		c.Redirect(http.StatusMovedPermanently, "/dashboardadmin")

		return
	}

	// JIKA USER BIASA
	c.String(http.StatusOK, "Login user berhasil")
}

func RegisterPage(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})

}

func Register(c *gin.Context) {

	nama := c.PostForm("nama")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := models.User{
		Nama:     nama,
		Email:    email,
		Password: password,
		Role:     "user",
	}

	config.DB.Create(&user)

	c.Redirect(http.StatusMovedPermanently, "/login")

}
func DashboardPage(c *gin.Context) {

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "Dashboard Admin",
	})
}
