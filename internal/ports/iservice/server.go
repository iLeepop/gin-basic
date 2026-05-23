package iservice

type IServerService interface {
	HealthCheck() (string, error)
}
