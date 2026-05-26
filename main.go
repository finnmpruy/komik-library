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

	// DEBUG WORKDIR (penting untuk kasus kamu)
	dir, _ := os.Getwd()
	fmt.Println("WORKDIR:", dir)

	// CONNECT DB
	config.ConnectDatabase()
	fmt.Println("Database berhasil terkoneksi")

	// MIGRATE
	config.DB.AutoMigrate(
		&models.Buku{},
		&models.User{},
	)

	// ENGINE
	r := gin.Default()

	// LOAD TEMPLATE
	r.LoadHTMLFiles(
		"templates/auth/login.html",
		"templates/auth/register.html",
		"templates/admin/dashboard.html",

		// File di dalam sub-folder buku
		"templates/admin/buku/index.html",
		"templates/admin/buku/create.html",
		"templates/admin/buku/edit.html",
	)

	// STATIC FILE
	r.Static("/static", "./static")

	// ROUTES
	routes.SetupRoutes(r)

	// RUN
	r.Run(":8080")
}
