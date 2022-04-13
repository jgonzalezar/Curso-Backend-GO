package routers

import(
	"errors"
	// "strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jgonzalezar/Curso-Backend-GO/bd"
	"github.com/jgonzalezar/Curso-Backend-GO/models"
)

/*Email valor de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string
/*ProcesoToken procesa el token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook") 
	claims := &models.Claim{}

	// splitToken := strings.Split(tk, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
	// if len(splitToken) != 2 {
	// 	return claims, false, string(""), errors.New("formato de Token invalido")
	// }

	// tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true{
			Email = claims.Email
			IDUsuario =  claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid{
		return claims, false, string(""), errors.New(" token invalido papa")
	}

	return claims, false, string(""), err
}