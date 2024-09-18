package config

import "net"

const grpcPort = "3000"

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func NewGRPCConfig() (GRPCConfig, error) {
	return &grpcConfig{
		port: grpcPort,
	}, nil
}
