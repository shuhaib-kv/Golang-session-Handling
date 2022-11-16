package main

import (
	"github.com/gin-gonic/gin"

	"work/controller"
	"work/db"
)

var R = gin.Default()

func init() {
	R.LoadHTMLGlob("templates/*.html")
}

func main() {
	db.Connect()

	R.GET("/", controller.Home)
	R.GET("/login", controller.Loginform)
	R.GET("/signup", controller.Signup)
	R.GET("/loginhandler", controller.Loginhandler)
	R.POST("/loginhandler", controller.Loginhandler)
	R.GET("/logoutu", controller.LogoutUser)
	R.POST("/signuphandler", controller.Signuphandler)
	R.GET("/userhome", controller.Userh)

	R.GET("/admin", controller.Admin)
	R.POST("/al", controller.AdminLoginHandler)
	R.GET("/ah", controller.AdminHome)
	R.GET("/block/:id", controller.Block)
	R.GET("/unblock/:id", controller.Unblock)
	R.GET("/adminlogout", controller.AdminLogout)
	R.GET("/um", controller.UserManagement)
	R.Run(":8080")
}
