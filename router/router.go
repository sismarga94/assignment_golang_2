package router

import (
	"assignment2/controllers"
	"assignment2/services"
	"database/sql"

	"github.com/antonlindstrom/pgstore"
	"github.com/gin-gonic/gin"
)

type Router struct {
	port  string
	db    *sql.DB
	store *pgstore.PGStore
}

func NewRouter(port string, db *sql.DB, store *pgstore.PGStore) *Router {
	return &Router{
		port:  port,
		db:    db,
		store: store,
	}
}

func (r *Router) Start() {
	router := gin.Default()

	controllers.AttachRouter(&controllers.HandlerConfig{
		R:           router,
		AuthService: services.NewAuthService(r.db),
		Store:       r.store,
	})
	router.Use(gin.Recovery())

	router.Run(r.port)
}
