package modelo

type Piloto struct {
	id       string
	licencia Licencia
}

func NuevoPiloto(id string, lic Licencia) (Piloto, error) {
	if id == "" {
		return Piloto{}, ErrIDPilotoVacio
	}

	return Piloto{id: id, licencia: lic}, nil
}
