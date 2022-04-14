package models

/*Tweet captura del body el mensaje que llega*/
type Tweet struct{
	Mensaje string `bson:"mensaje" json:"mensaje"`
}