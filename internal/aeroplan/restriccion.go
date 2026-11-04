package aeroplan

type Restriccion struct {
	alturaMaxima        float64
	tipo                string
	licenciasRequeridas []Licencia
}
