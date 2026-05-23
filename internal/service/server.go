package server_impl

type ServerService struct{}

func (s *ServerService) HealthCheck() (string, error) {
	return "ok", nil
}
