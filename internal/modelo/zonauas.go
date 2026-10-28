package modelo

// Su id lo otorga los datos de ENAIRE a través de sus archivos descargables "ZONAS UAS".
type ZonaUAS struct {
	id            string
	geometria     Geometria
	restricciones []Restriccion
}

func NuevaZonaUAS(id string, geo Geometria, rest []Restriccion) (ZonaUAS, error) {
	if id == "" {
		return ZonaUAS{}, ErrIDEnaireVacio
	}
	return ZonaUAS{
		id:            id,
		geometria:     geo,
		restricciones: rest,
	}, nil
}
