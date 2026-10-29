package modelo

import "errors"

var (
	ErrCategoriaNoReconocida = errors.New("categoría de licencia no reconocida")
	ErrCoordenadasInvalidas  = errors.New("coordenadas geográficas fuera de rango o inválidas")
	ErrIDPilotoVacio         = errors.New("el ID del piloto no puede estar vacío")
	ErrIDEnaireVacio         = errors.New("el identificador de ENAIRE es obligatorio")
	ErrAlturaNegativa        = errors.New("la altura no puede ser negativa")
	ErrIDEncargoVacio        = errors.New("el ID del encargo no puede estar vacío")
	ErrTipoRestriccionVacio  = errors.New("el tipo de restricción no puede estar vacío")
	ErrGeometriaInvalida     = errors.New("una geometría necesita al menos 3 puntos para formar un polígono")
	ErrZonaDuplicada         = errors.New("ya existe una zona con este ID")
)
