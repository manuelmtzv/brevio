package models

type HealthStatus string

const (
	HealthOK       HealthStatus = "ok"
	HealthDegraded HealthStatus = "degraded"
	HealthDown     HealthStatus = "down"
)

type HealthCheck struct {
	Status HealthStatus `json:"status"`
}
