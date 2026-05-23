package server

type ServerService struct{}

func (s *ServerService) HealthCheck() (string, error) {
	return "ok", nil
}
