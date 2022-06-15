package main

import (
	"fmt"
	"forum/controllers"
	DBM "forum/repository"
	"forum/tools"
	"forum/tools/riot"
	_ "forum/tools/session"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	err := tools.LoadEnv(".env")
	rand.Seed(time.Now().Unix())

	riot.API.SetKey(os.Getenv("riot_key"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	DBM.InitializeDatabase("forum_database.db")
	http.Handle("/", &controllers.ClientController{})
	http.Handle("/api/", &controllers.APIController{})
	fmt.Println("Start...")
	fmt.Println("\thttp://localhost:8080")
	tools.Openbrowser("http://localhost:8080")
	//static
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./src/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./src/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./src/js"))))

	http.ListenAndServe(":8080", nil)
}
