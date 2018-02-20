package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/coopernurse/gorp"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
)

var dbmap *gorp.DbMap

// Init ()
func init() {

	db, err := sql.Open("sqlite3", "db.bin")
	if err != nil {
		log.Fatalln("Fail to create database", err)
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "ID")
	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatalln("Could not build tables", err)
	}

	// user := MyUserModel{1, "testuser", "password", false}
	// err = dbmap.Insert(&user)
	// if err != nil {
	// 	log.Fatalln("Could not insert test user", err)
	// }

	// dbmap
	// return dbmap
}

func main() {

	m := martini.Classic()
	// dbmap = initDb()

	m.Use(gzip.All(gzip.Options{
		CompressionLevel: gzip.BestCompression,
	}))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layouts/layout",
		Extensions: []string{".html"},
		// Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Delims:     render.Delims{Left: "[[", Right: "]]"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(http.StatusOK, "index", nil)
	})

	m.Get("/register", func(r render.Render) {
		r.HTML(http.StatusOK, "register", nil)
	})

	m.Post("/register", binding.Bind(User{}), func(user User) string {
		return fmt.Sprintf("Name: %s\nPassword: %s\nPassword2: %s\n",
			user.Username, user.Password, user.Password2)
	})

	m.Get("/login", func(r render.Render) {
		r.HTML(http.StatusOK, "login", nil)
	})

	m.Post("/login", binding.Bind(User{}), func(user User) string {
		return fmt.Sprintf("Name: %s\nPassword: %s\n",
			user.Username, user.Password)
	})

	m.Run()
}
