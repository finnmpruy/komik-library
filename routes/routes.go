package routes

import (
	"komik-library/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/", controllers.LoginPage)

	r.GET("/login", controllers.LoginPage)

	r.POST("/login", controllers.Login)

	r.GET("/register", controllers.RegisterPage)

	r.POST("/register", controllers.Register)

	r.GET("/dashboardadmin", controllers.DashboardPage)

	r.GET("/", controllers.Home)

	r.GET("/buku", controllers.IndexBuku)
	r.GET("/buku/create", controllers.CreateBuku)
	r.POST("/buku/store", controllers.StoreBuku)

	r.GET("/buku/show/:id", controllers.ShowBuku)

	r.GET("/buku/edit/:id", controllers.EditBuku)
	r.POST("/buku/update/:id", controllers.UpdateBuku)

	r.GET("/buku/delete/:id", controllers.DeleteBuku)

}
