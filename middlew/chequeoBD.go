package middlew

import(
	"net/http"
	"github.com/jgonzalezar/Curso-Backend-GO/bd"
	
)

/*ChequeoBD es el middleware que permite conocer el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if bd.ChequeoConnection()==0{
			http.Error(w, "conexión perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w,r)
	}
}