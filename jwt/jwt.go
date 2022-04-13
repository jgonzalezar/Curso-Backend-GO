package jwt

import(
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jgonzalezar/Curso-Backend-GO/models"
)

/*GeneroJWT genera el encriptado con jwt*/
func GeneroJWT(t models.Usuario) (string, error){
	miClave := []byte("JuanilloPillo_Carajillo")

	payload := jwt.MapClaims{
		"email": 		t.Email,
		"nombre":		t.Nombre,
		"apellidos": 	t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": 	t.Biografia,
		"ubicacion": 	t.Ubicacion,
		"sitioweb":		t.SitioWeb,
		"_id": 			t.ID.Hex(),
		"exp": time.Now().Add(24*time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil{
		return tokenStr, err
	}
	return tokenStr, nil
}