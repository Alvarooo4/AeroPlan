package modelo

type CatalogoZonas struct {
	zonas map[string]ZonaUAS
}

func NuevoCatalogoZonas() CatalogoZonas {
	return CatalogoZonas{
		zonas: make(map[string]ZonaUAS),
	}
}

func (c *CatalogoZonas) Registrar(id string, geometria Geometria, restricciones []Restriccion) error {
	if _, existe := c.zonas[id]; existe {
		return ErrZonaDuplicada
	}

	zona, err := nuevaZonaUAS(id, geometria, restricciones)
	if err != nil {
		return err
	}

	c.zonas[id] = zona
	return nil
}
