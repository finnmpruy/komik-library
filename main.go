package main

import (
	"fmt"
	"os"

	"komik-library/config"
	"komik-library/models"
	"komik-library/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// DEBUG WORKDIR
	dir, _ := os.Getwd()
	fmt.Println("WORKDIR:", dir)

	// CONNECT DB
	config.ConnectDatabase()
	fmt.Println("Database berhasil terkoneksi")

	// MIGRATE
	config.DB.AutoMigrate(
		&models.Buku{},
		&models.User{},
		&models.Peminjaman{},
	)

	// ENGINE
	r := gin.Default()

	// LOAD TEMPLATE
	r.LoadHTMLFiles(
		"templates/auth/login.html",
		"templates/auth/register.html",
		"templates/admin/dashboard.html",
		"templates/admin/buku/index.html",
		"templates/admin/buku/create.html",
		"templates/admin/buku/edit.html", // Baris duplikat sudah dihapus
		"templates/admin/peminjaman/index_peminjaman.html",
		"templates/user/dashboard_user.html",
	)

	// ====== BAGIAN STATIC FILES DISATUKAN DI SINI ======
	r.Static("/static", "./static")

	r.Static("/assets/uploads", "./static/uploads")

	// ROUTES
	routes.SetupRoutes(r)

	// RUN
	r.Run(":8080")
}
