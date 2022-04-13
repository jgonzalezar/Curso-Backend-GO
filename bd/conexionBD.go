package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BD*/
var MongoCN = ConectarBD()
var clienOptions = options.Client().ApplyURI("mongodb+srv://jgonzalezar:4NblBFqrn3Id4R5k@curso-backend-go.l9zw3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarBD es la función que permite conectarse a la base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
