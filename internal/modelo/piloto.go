package modelo

import "errors"

type Piloto struct {
	id          string
	licencia Licencia
}

func NuevoPiloto(id string, cert Licencia) (Piloto, error) {
	if id == "" {
		return Piloto{}, errors.New("el ID del piloto no puede estar vacío")
	}

	return Piloto{id: id, certificado: cert}, nil
}