package aeroplan

type CatalogoZonas struct {
	zonas map[string]ZonaUAS
}

func NewCatalogoZonas() CatalogoZonas {
	return CatalogoZonas{
		zonas: make(map[string]ZonaUAS, 0),
	}
}
