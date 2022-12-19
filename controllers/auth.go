package controllers

import (
	"assignment2/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Web",
	})
}

func (h *Handler) HandleRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Web",
	})
}

func (h *Handler) HandleRegister(c *gin.Context) {
	var data dto.RegisterDto = dto.RegisterDto{
		Username:  c.PostForm("username"),
		Password:  c.PostForm("password"),
		Firstname: c.PostForm("firstname"),
		Lastname:  c.PostForm("lastname"),
	}

	fmt.Printf("data register :%v\n", data)
	_, err := h.AuthService.UserRegister(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bad Request")
	}
	c.HTML(http.StatusOK, "register_success.html", gin.H{
		"title": "Web",
	})
}

func (h *Handler) HandleLogin(c *gin.Context) {
	session, _ := h.Store.Get(c.Request, "SESSION_ID")
	session.Values["username"] = c.PostForm("username")

	//store session
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var data dto.LoginDto = dto.LoginDto{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
	}

	fmt.Printf("data register :%v\n", data)
	_, err = h.AuthService.Login(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/home")
}

func (h *Handler) HandleHome(c *gin.Context) {
	session, _ := h.Store.Get(c.Request, "SESSION_ID")
	if len(session.Values) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":  "no sessions",
		})
		return
	}
	var username string = session.Values["username"].(string)
	user, err := h.AuthService.GetUser(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"username":  user.Username,
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
		"status":    "Success",
	})

}
