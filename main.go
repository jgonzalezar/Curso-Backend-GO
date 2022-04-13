package main

import (
	"log"
	"github.com/jgonzalezar/Curso-Backend-GO/handlers"
	"github.com/jgonzalezar/Curso-Backend-GO/bd"

)

func main(){
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexion a BD")
		return
	}
	handlers.Manejadores()
}