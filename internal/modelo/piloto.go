package modelo

import "errors"

type Piloto struct {
	id          string
	certificado Certificado
}

func NuevoPiloto(id string, cert Certificado) (Piloto, error) {
	if id == "" {
		return Piloto{}, errors.New("el ID del piloto no puede estar vacío")
	}
	
	return Piloto{id: id, certificado: cert}, nil
}