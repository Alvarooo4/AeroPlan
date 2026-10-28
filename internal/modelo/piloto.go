package modelo

// Piloto representa a un operador de drones en el sistema.
// Su ID lo otorga AESA (Agencia Estatal de Seguridad Aérea) una vez superado el examen de piloto.
// Sigue la estructura "ESP-RP-000000xxxxxx", donde las 6 "x" son el id propio de cada piloto.
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
