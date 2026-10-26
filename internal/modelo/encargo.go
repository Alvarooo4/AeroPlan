package modelo

import "errors"

type Encargo struct {
	id           string
	localizacion Localizacion
}

func NuevoEncargo(id string, loc Localizacion) (Encargo, error) {
	if id == "" {
		return Encargo{}, errors.New("el ID del encargo es obligatorio")
	}
	
	return Encargo{id: id, localizacion: loc}, nil
}