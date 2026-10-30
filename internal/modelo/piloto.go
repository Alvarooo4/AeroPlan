package modelo

type Piloto struct {
	id       string
	licencia Licencia
}

func NuevoPiloto(id string, lic Licencia) (Piloto, error) {
	if id == "" {
		return Piloto{}, ErrorIDPilotoVacio{}
	}

	return Piloto{id: id, licencia: lic}, nil
}
