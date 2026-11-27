package aeroplan

type PlanDeVuelo struct {
	encargosPorUbicacion map[Localizacion]Encargo
	zonasUAS             []ZonaUAS
}
