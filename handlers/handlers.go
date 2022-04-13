package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/jgonzalezar/Curso-Backend-GO/middlew"
	"github.com/jgonzalezar/Curso-Backend-GO/routers"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	
)

/*Manejadores seteo el puerto, el Handler y pongo a escuchar al servidor*/
func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")




	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}