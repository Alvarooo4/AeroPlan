package modelo

type Restriccion struct {
	alturaMaxima float64
	tipo         string
}

func NuevaRestriccion(altura float64, tipo string) (Restriccion, error) {
	if altura < 0 {
		return Restriccion{}, ErrAlturaNegativa
	}

	if tipo == "" {
		return Restriccion{}, ErrTipoRestriccionVacio
	}

	return Restriccion{alturaMaxima: altura, tipo: tipo}, nil
}
