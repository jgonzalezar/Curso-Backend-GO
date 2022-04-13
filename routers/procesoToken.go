package routers

import(
	"errors"
	"strings"

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
	miClave := []byte("JuanilloPillo_Carajillo") 
	claims := &models.Claim{}

	spliToken := strings.Split(tk, "Bearer")
	if len(spliToken) != 2 {
		return claims, false, string(""), errors.New("formato de Token invalido")
	}

	tk = strings.TrimSpace(spliToken[1])

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
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}