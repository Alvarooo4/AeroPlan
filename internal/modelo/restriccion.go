package modelo

import "errors"

type Restriccion struct {
	alturaMaxima float64
	tipo         string
}

func NuevaRestriccion(altura float64, tipo string) (Restriccion, error) {
	if altura < 0 {
		return Restriccion{}, errors.New("la altura máxima no puede ser negativa")
	}
	
	if tipo == "" {
		return Restriccion{}, errors.New("el tipo de restricción es obligatorio")
	}

	return Restriccion{alturaMaxima: altura, tipo: tipo}, nil
}