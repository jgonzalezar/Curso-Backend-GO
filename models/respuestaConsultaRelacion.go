package models

/*RespuestaConsultaRelacion tiene el true o false que se obtiene de consultar la relación emtre 2 usuarios*/
type RespuestaConsultaRelacion struct{
	Status bool `json:"status"`
}