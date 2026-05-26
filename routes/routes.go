package routes

import (
	"komik-library/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// AUTH
	r.GET("/", controllers.LoginPage)
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.Login)

	r.GET("/register", controllers.RegisterPage)
	r.POST("/register", controllers.Register)

	// DASHBOARD ADMIN
	r.GET("/dashboardadmin", controllers.DashboardPage)

	//DASHBOARD USER
	r.GET("/dashboarduser", controllers.DashboardUser)

	// GRUP ADMIN
	admin := r.Group("/admin")
	{
		// sub-grup kelola buku
		buku := admin.Group("/buku")
		{
			buku.GET("/", controllers.IndexBuku)
			buku.GET("/create", controllers.CreateBuku)
			buku.POST("/store", controllers.StoreBuku)
			buku.GET("/edit/:id", controllers.EditBuku)
			buku.POST("/update/:id", controllers.UpdateBuku)
			buku.GET("/delete/:id", controllers.DeleteBuku)
		}

		peminjaman := admin.Group("/peminjaman")
		{
			peminjaman.GET("/", controllers.IndexPeminjaman)

		}
	}
}
