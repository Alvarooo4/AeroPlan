package modelo

import "errors"

type Licencia struct {
	categoria string
}

func NuevaLicencia(categoria string) (Licencia, error) {
	if categoria != "A1/A3" && categoria != "A2" && categoria != "STS" {
		return Licencia{}, errors.New("categoría de licencia no reconocida")
	}
	
	return Licencia{categoria: categoria}, nil
}