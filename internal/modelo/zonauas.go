package modelo

import "errors"

type ZonaUAS struct {
	idEnaire      string
	restricciones []Restriccion
}

func NuevaZonaUAS(idEnaire string, restricciones []Restriccion) (ZonaUAS, error) {
	if idEnaire == "" {
		return ZonaUAS{}, errors.New("el identificador de ENAIRE es obligatorio")
	}
	
	return ZonaUAS{idEnaire: idEnaire, restricciones: restricciones}, nil
}