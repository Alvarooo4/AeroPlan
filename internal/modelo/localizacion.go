package modelo

import "errors"

type Localizacion struct {
	latitud  float64
	longitud float64
}

func NuevaLocalizacion(lat, lon float64) (Localizacion, error) {
	if lat < -90 || lat > 90 || lon < -180 || lon > 180 {
		return Localizacion{}, errors.New("coordenadas geográficas fuera de rango")
	}
	
	return Localizacion{latitud: lat, longitud: lon}, nil
}