package modelo

type ZonaUAS struct {
	idEnaire            string
	restricciones       []Restriccion
	geometria           string // Extensión geográfica (polígono, círculo, etc.)
	licenciasPermitidas []Licencia
}

func NuevaZonaUAS(idEnaire string, restricciones []Restriccion, geometria string, licencias []Licencia) (ZonaUAS, error) {
	if idEnaire == "" {
		return ZonaUAS{}, ErrIDEnaireVacio
	}

	return ZonaUAS{
		idEnaire:            idEnaire,
		restricciones:       restricciones,
		geometria:           geometria,
		licenciasPermitidas: licencias,
	}, nil
}
