package modelo

import "math"

type Localizacion struct {
	latitud  float64
	longitud float64
}

func NuevaLocalizacion(lat, lon float64) (Localizacion, error) {
	if math.IsNaN(lat) || math.IsNaN(lon) || lat < -90 || lat > 90 || lon < -180 || lon > 180 {
		return Localizacion{}, ErrorCoordenadasInvalidas{Motivo: "valores fuera de rango o NaN"}
	}

	return Localizacion{latitud: lat, longitud: lon}, nil
}
