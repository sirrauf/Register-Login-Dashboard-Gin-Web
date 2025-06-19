package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "dashboardapp/model"
)

func main() {
	InitDB()
	routers := gin.Default()
	routers.LoadHTMLGlob("templates/*")
	routers.Static("/static", "./static")

	routers.GET("/", func(content *gin.Context) {
		content.HTML(http.StatusOK, "register.html", gin.H{"title": "Register"})
	})

	routers.POST("/register", func(content *gin.Context) {
		var user User
		user.NamaLengkap = content.PostForm("nama_lengkap")
		user.Alamat = content.PostForm("alamat")
		user.Email = content.PostForm("email")
		user.Password = content.PostForm("password")
		user.Role = content.PostForm("role")

		if errors := Database.Create(&user).Error; errors != nil {
			content.String(http.StatusInternalServerError, "Gagal register: %v", errors)
			return
		}

		content.Redirect(http.StatusFound, "/login")
	})

	routers.GET("/login", func(content *gin.Context) {
		content.HTML(http.StatusOK, "login.html", gin.H{"title": "Login"})
	})

	routers.POST("/login", func(content *gin.Context) {
		email := content.PostForm("email")
		password := content.PostForm("password")

		var user User
		if errors := Database.Where("email = ? AND password = ?", email, password).First(&user).Error; errors != nil {
			content.String(http.StatusUnauthorized, "Login gagal!")
			return
		}

		content.HTML(http.StatusOK, "dashboard.html", gin.H{
			"title": "Dashboard",
			"nama":  user.NamaLengkap,
			"email": user.Email,
			"role":  user.Role,
		})
	})

	routers.Run(":8080")
}
