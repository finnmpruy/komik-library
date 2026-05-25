package main

import (
	"komik-library/config"
	"komik-library/models"
	"komik-library/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&models.Buku{},
		&models.User{},
	)

	r := gin.Default()

	// LOAD TEMPLATE
	r.LoadHTMLFiles(
		"templates/auth/login.html",
		"templates/auth/register.html",
		"templates/admin/dashboard.html",
	)
	// STATIC FILE
	r.Static("/static", "./static")

	// UPLOADS

	r.Static("/uploads", "./uploads")

	// ROUTES
	routes.SetupRoutes(r)

	// RUN SERVER
	r.Run(":8080")
}
