package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type Env struct {
	db *sqlx.DB
}

type App struct {
	db Datastore
}



func main() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PW"), os.Getenv("DB_INSTANCE"), os.Getenv("DB_NAME"))
	db := NewDB(datasource)
	//env := &Env{db}

	app := &App{db}

	router := NewRouter(env)

	port := ":5000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	log.Fatal(http.ListenAndServe(port, router))
}