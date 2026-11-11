package aeroplan

type TipoRestriccion string

const (
	RestriccionProhibida      TipoRestriccion = "PROHIBIDA"
	RestriccionRequisitos     TipoRestriccion = "REQ_AUTHORIZATION"
	RestriccionControlada     TipoRestriccion = "CTR"
	RestriccionLibre          TipoRestriccion = "LIBRE"
)

const AlturaMaximaLegal float64 = 120.0

type Restriccion struct {
	alturaMaximaZona float64
	tipo             TipoRestriccion
}
