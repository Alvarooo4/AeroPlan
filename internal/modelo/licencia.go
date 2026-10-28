package modelo

type Licencia string

const (
	A1A3 Licencia = "A1/A3"
	A2   Licencia = "A2"
	STS  Licencia = "STS"
)

func NuevaLicencia(categoria string) (Licencia, error) {
	lic := Licencia(categoria)

	if lic != A1A3 && lic != A2 && lic != STS {
		return "", ErrCategoriaNoReconocida
	}

	return lic, nil
}
