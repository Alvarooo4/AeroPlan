package modelo

type Geometria struct {
	puntos []Localizacion
}

func NuevaGeometria(puntos []Localizacion) (Geometria, error) {
	if len(puntos) < 3 {
		return Geometria{}, ErrGeometriaInvalida
	}
	return Geometria{puntos: puntos}, nil
}
