package main

import (
	"GenDriver/src/routes"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.SetPrefix("Ocorreu um error: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

}
func main() {
	// db := database.ConectarDB()
	// fmt.Println(db)
	routes.Router()
}
