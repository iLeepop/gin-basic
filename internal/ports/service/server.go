package service

type IServerService interface {
	HealthCheck() (string, error)
}
