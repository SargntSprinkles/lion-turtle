package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SargntSprinkles/lion-turtle/characters"
	"github.com/SargntSprinkles/lion-turtle/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jinzhu/gorm"
)

var env map[string]string
var lionTurtleDB *gorm.DB

func main() {
	getEnv()

	for _, v := range []string{"DBHOST", "DBPORT", "DBUSER", "DBPW", "DBDB"} {
		if len(env[v]) == 0 {
			log.Print("Need credentials to connect to the database")
			return
		}
	}

	connectDB()
	loadData()

	log.Print("Setting up routes")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", index)
	r.Route("/character", characterRouter)
	r.Route("/reference", referenceRouter)

	log.Printf("Listening on port %s", env["PORT"])
	log.Fatal(http.ListenAndServe(":"+env["PORT"], r))
}

func getEnv() {
	env = map[string]string{}
	env["PORT"] = setEnv("PORT", "8080")
	env["DBHOST"] = setEnv("DBHOST", "")
	env["DBPORT"] = setEnv("DBPORT", "3306")
	env["DBUSER"] = setEnv("DBUSER", "")
	env["DBPW"] = setEnv("DBPW", "")
	env["DBDB"] = setEnv("DBDB", "lionturtle")
}

func setEnv(envName, defaultVal string) string {
	newEnv := os.Getenv(envName)
	if newEnv == "" && defaultVal != "" {
		return defaultVal
	}
	return newEnv
}

func connectDB() {
	log.Print("Connecting to DB")
	dbConfig := db.Config{
		Host:     env["DBHOST"],
		Port:     env["DBPORT"],
		User:     env["DBUSER"],
		Password: env["DBPW"],
		Database: env["DBDB"],
	}
	dbConn := db.New(dbConfig)
	var dbErr error
	lionTurtleDB, dbErr = dbConn.Connect()
	if dbErr != nil {
		log.Fatalf("Error connecting to database: %s", dbErr.Error())
	}
	characters.LionTurtleDB = lionTurtleDB
	log.Print("Connected!")
}

func loadData() {
	// just in case?
}

func newTemplate(files ...string) *template.Template {
	files = append(files, "templates/base.gohtml", "templates/footer.gohtml", "templates/style.gohtml")
	return template.Must(template.ParseFiles(files...))
}

func index(w http.ResponseWriter, r *http.Request) {
	indexTemplate := newTemplate("templates/index.gohtml")
	myCharacters := characters.GetAllCharacters()
	viewData := map[string]interface{}{}
	viewData["characters"] = myCharacters
	viewData["highlight"] = "character"
	indexTemplate.ExecuteTemplate(w, "base", viewData)
}
