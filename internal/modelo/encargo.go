package modelo

// Encargo representa un trabajo de vuelo a realizar.
// Su ID utiliza la nomenclatura interna de la empresa ZanganoPictures con la estructura "[Año]E-x".
// El "[Año]" corresponde al año de realización, "E" es una letra identificativa inmutable, y "x" es el número incremental del encargo (ej. 2026E-5).
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
