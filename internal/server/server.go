package server

import (
	"fmt"

	"go.uber.org/zap"
)

func (s Server) Run() error {
	var err error

	r, err := s.NewRouter()
	if err != nil {
		return err
	}

	s.logger.Debug("Listen and Server",
		zap.String("IP", s.config.BindIp()),
		zap.Int("Port", s.config.BindPort()))
	err = r.Run(
		fmt.Sprintf("%s:%d", s.config.BindIp(), s.config.BindPort()),
	)

	return err
}
