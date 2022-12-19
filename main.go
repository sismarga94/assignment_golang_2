package main

import (
	"assignment2/db"
	"assignment2/router"
)

func main() {
	dbPostgres, _ := db.NewPostgres()
	store, _ := db.NewPostgresStore(dbPostgres)
	store.Options.MaxAge = 300

	r := router.NewRouter(":9999", dbPostgres, store)
	r.Start()
}
