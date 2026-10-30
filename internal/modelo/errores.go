package modelo

import "fmt"

type ErrorCategoriaNoReconocida struct {
	Categoria string
}

func (e ErrorCategoriaNoReconocida) Error() string {
	return fmt.Sprintf("categoría de licencia no reconocida: %s", e.Categoria)
}

type ErrorCoordenadasInvalidas struct {
	Motivo string
}

func (e ErrorCoordenadasInvalidas) Error() string {
	return fmt.Sprintf("coordenadas geográficas fuera de rango o inválidas: %s", e.Motivo)
}

type ErrorIDPilotoVacio struct{}

func (e ErrorIDPilotoVacio) Error() string {
	return "el ID del piloto no puede estar vacío"
}

type ErrorIDEnaireVacio struct{}

func (e ErrorIDEnaireVacio) Error() string {
	return "el identificador de ENAIRE es obligatorio"
}

type ErrorAlturaNegativa struct {
	Altura float64
}

func (e ErrorAlturaNegativa) Error() string {
	return fmt.Sprintf("la altura no puede ser negativa: %f", e.Altura)
}

type ErrorIDEncargoVacio struct{}

func (e ErrorIDEncargoVacio) Error() string {
	return "el ID del encargo no puede estar vacío"
}

type ErrorTipoRestriccionVacio struct{}

func (e ErrorTipoRestriccionVacio) Error() string {
	return "el tipo de restricción no puede estar vacío"
}

type ErrorGeometriaInvalida struct {
	Puntos int
}

func (e ErrorGeometriaInvalida) Error() string {
	return fmt.Sprintf("una geometría necesita al menos 3 puntos para formar un polígono, puntos recibidos: %d", e.Puntos)
}

type ErrorZonaDuplicada struct {
	ID string
}

func (e ErrorZonaDuplicada) Error() string {
	return fmt.Sprintf("ya existe una zona con este ID: %s", e.ID)
}
