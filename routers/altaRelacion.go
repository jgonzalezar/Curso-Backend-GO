package routers

import(
	"net/http"
	"github.com/jgonzalezar/Curso-Backend-GO/bd"
	"github.com/jgonzalezar/Curso-Backend-GO/models"
)

/*AltaRelacion realiza el registro de la relación entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "El parametro id es obligatrio",http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al insertar relación" + err.Error(),http.StatusBadRequest)
		return
	}
	if status == false{
		http.Error(w, "No se ha logrado insertar la relación",http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}