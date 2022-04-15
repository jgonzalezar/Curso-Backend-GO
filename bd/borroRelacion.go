package bd

import(
	"context"
	"time"

	"github.com/jgonzalezar/Curso-Backend-GO/models"
)

/*BorroRelacion borra la relacion de la BD*/
func BorroRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}