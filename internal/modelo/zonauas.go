package modelo

type ZonaUAS struct {
	id            string
	geometria     Geometria
	restricciones []Restriccion
}

func nuevaZonaUAS(id string, geo Geometria, rest []Restriccion) (ZonaUAS, error) {
	if id == "" {
		return ZonaUAS{}, ErrorIDEnaireVacio{}
	}
	return ZonaUAS{
		id:            id,
		geometria:     geo,
		restricciones: rest,
	}, nil
}
