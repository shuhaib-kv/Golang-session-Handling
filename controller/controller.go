package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"tawesoft.co.uk/go/dialog"

	"work/db"
	"work/models"
)

func Loginform(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.HTML(http.StatusOK, "login.html", nil)

}
func Home(C *gin.Context) {
	C.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	C.Writer.Header().Set("Pragma", "no-cache")
	C.Writer.Header().Set("Expires", "0")
	ok := UserLoged(C)
	if ok {
		C.Redirect(303, "/userhome")
		return
	}
	C.HTML(http.StatusOK, "home.html", nil)
}

func Login(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	ok := UserLoged(c)
	if ok {
		c.Redirect(303, "/userhome")
		return
	}
	c.HTML(http.StatusOK, "/login", nil)

}

func UserLoged(c *gin.Context) bool {
	session, _ := Store.Get(c.Request, "session")
	_, ok := session.Values["email"]
	if !ok {
		return ok
	}
	return true
}
func Loginhandler(C *gin.Context) {
	C.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	err := C.Request.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
	}
	Emails := C.PostForm("email")
	password := C.PostForm("password")
	var user models.User
	var status bool
	db.DB.Raw("select email,password,block_status FROM users where email=?", Emails).Scan(&user)
	if user.Email == Emails && user.Password == password && user.Block_status == false {

		status = true
	}
	if !status {
		if user.Block_status == true {
			dialog.Alert("hey you are blocked by Admin")
			C.Redirect(303, "/login")
			return
		} else {
			dialog.Alert("wrong username and password")
			C.Redirect(303, "/login")
			return
		}
	}
	sessions, err := Store.Get(C.Request, "session")
	sessions.Values["email"] = Emails
	P = sessions.Values["email"]
	sessions.Save(C.Request, C.Writer)
	C.Redirect(301, "/userhome")

}
func Userh(C *gin.Context) {
	ok := Userloggedin(C)
	if !ok {
		C.Redirect(303, "/")
		return
	}

	C.HTML(200, "userhome.html", gin.H{})

}

func Signup(C *gin.Context) {
	C.HTML(200, "signup.html", nil)
}
func Signuphandler(C *gin.Context) {
	err := C.Request.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
	}
	name := C.PostForm("name_si")
	email := C.PostForm("email_si")
	password := C.PostForm("password_si")
	user := models.User{Name: name, Email: email, Password: password}
	db.DB.Create(&user)
	dialog.Alert("Hey %s, Your account is successfully created. Click OK to LOGIN!", user.Name)
	C.Redirect(http.StatusSeeOther, "/login")
}

func LogoutUser(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	sessions, err := Store.Get(c.Request, "session")
	sessions.Options.MaxAge = -1
	sessions.Save(c.Request, c.Writer)
	fmt.Println(err)
	c.Redirect(303, "/")
}

func Userloggedin(g *gin.Context) bool {
	g.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	session, err := Store.Get(g.Request, "session")
	if session.Values["email"] == nil {
		return false
	}
	fmt.Println(err)
	return true
}

var Store = sessions.NewCookieStore([]byte("admin"))
var P interface{}

func init() {

}
func HomePage(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	session, _ := Store.Get(c.Request, "session")
	String := session.Values["email"]
	h := String.(string)
	splt := strings.TrimSuffix(h, "@gmail.com")
	ok := Userloggedin(c)
	if !ok {
		c.Redirect(303, "/")
		return
	}
	c.HTML(200, "homepage.html", gin.H{
		"name": splt,
	})

}

var AdminDB = map[string]string{
	"password": "admin",
	"email":    "admin@gmail.com",
}

func LogoutAdmin(c *gin.Context) {

	cookie, err := c.Request.Cookie("adminsession")
	if err != nil {
		c.Redirect(303, "/admin")
	}
	c.SetCookie("adminsession", "", -1, "/", "localhost", false, false)
	_ = cookie
	c.Redirect(http.StatusSeeOther, "/admin")

}
func Admin(C *gin.Context) {
	ok := AdminLoged(C)
	if !ok {
		C.HTML(http.StatusOK, "adminlogin.html", nil)
		return
	}
	C.HTML(200, "adminhome.html", nil)
}
func AdminLoginHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println("error parsing form")
	}
	var statuss bool
	email := c.PostForm("aemail")
	password := c.PostForm("apassword")
	if AdminDB["email"] == email && AdminDB["password"] == password && c.Request.Method == "POST" {
		statuss = true
	}
	if !statuss {
		dialog.Alert("wrong user name and password")
		c.Redirect(303, "/admin")
	}
	session, err := Store.Get(c.Request, "admin")
	session.Values["email"] = "email"
	session.Save(c.Request, c.Writer)
	c.Redirect(301, "/ah")
}

func AdminLoged(c *gin.Context) bool {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	session, _ := Store.Get(c.Request, "admin")
	_, ok := session.Values["email"]
	if !ok {
		return ok
	}
	return true
}

type Page struct {
	Status bool
}

var Status Page

func AdminHome(C *gin.Context) {
	C.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ok := AdminLoged(C)
	if !ok {
		C.Redirect(303, "/admin")
		return
	}
	C.HTML(200, "adminhome.html", nil)
}
func UserManagement(C *gin.Context) {
	C.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ok := AdminLoged(C)
	if !ok {
		C.Redirect(303, "/admin")
		return
	}
	Status.Status = true
	// params := C.Param("id")
	// page, _ := strconv.Atoi(params)
	var user []models.User
	db.DB.Raw("SELECT * FROM users ").Scan(&user)
	// db.DB.Raw("SELECT * FROM users ORDER BY id ASC").Scan(&user)

	C.HTML(200, "usermanagement.html", gin.H{
		"user":   user,
		"status": Status.Status,
	})
}
func Block(c *gin.Context) {
	params := c.Param("id")
	page, _ := strconv.Atoi(params)
	var users models.User
	db.DB.Raw("update users SET block_status=true WHERE id=?", page).Scan(&users)
	c.Redirect(303, "/um")
}
func Unblock(c *gin.Context) {

	params := c.Param("id")
	page, _ := strconv.Atoi(params)
	var users models.User
	db.DB.Raw("update users SET block_status=false WHERE id=?", page).Scan(&users)
	c.Redirect(303, "/um")
}
func AdminLogout(c *gin.Context) { // adminLogout page
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	session, err := Store.Get(c.Request, "admin")
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	fmt.Println(err)
	c.Redirect(303, "/admin")

}
