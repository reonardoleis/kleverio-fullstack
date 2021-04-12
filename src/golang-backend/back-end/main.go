package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	utxo "github.com/reonardoleis/klever.io-backend/src/utxo/controllers"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/balance/{address}", utxo.GetBalance).Methods("GET")
	router.HandleFunc("/health", utxo.CheckHealth).Methods("GET")

	//Apenas para servir arquivo estático (site em ReactJS) na nuvem,
	//em caso de uso local pode ser desconsiderado (caso queira)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public"))).Methods("GET")
	port := getPort()
	log.Println("Server started on port " + port)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(port, handler))
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}
