package modelo

type Encargo struct {
	id           string
	localizacion Localizacion
}

func NuevoEncargo(id string, loc Localizacion) (Encargo, error) {
	if id == "" {
		return Encargo{}, ErrIDEncargoVacio
	}

	return Encargo{id: id, localizacion: loc}, nil
}
