package aeroplan

type IdZona string

type PlanDeVuelo struct {
	encargosPorUbicacion map[Localizacion]Encargo
	zonasUAS map[IdZona]ZonaUAS
}
