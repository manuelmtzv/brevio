package models

type HealthCheck struct {
	Status HealthStatus `json:"status"`
}

type HealthStatus string

const (
	OK      HealthStatus = "OK"
	Warning HealthStatus = "WARNING"
	Error   HealthStatus = "ERROR"
)
