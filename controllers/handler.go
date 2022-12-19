package controllers

import (
	"assignment2/services"

	"github.com/antonlindstrom/pgstore"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	AuthService services.AuthServiceProvider
	Store       *pgstore.PGStore
}

type HandlerConfig struct {
	R           *gin.Engine
	AuthService services.AuthServiceProvider
	Store       *pgstore.PGStore
}

func AttachRouter(c *HandlerConfig) {
	h := &Handler{
		AuthService: c.AuthService,
		Store:       c.Store,
	}

	c.R.Static("/static", "./static")
	c.R.LoadHTMLGlob("static/*.html")
	c.R.GET("/", h.HandleIndex)
	c.R.POST("/home", h.HandleHome)
	c.R.GET("/register", h.HandleRegisterPage)
	c.R.POST("/register", h.HandleRegister)
	c.R.POST("/login", h.HandleLogin)
}
