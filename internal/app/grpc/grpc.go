package app

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ZhdanovichVlad/service-podof/internal/config"
	grpccontroller "github.com/ZhdanovichVlad/service-podof/internal/controller/grpc"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	pb "github.com/ZhdanovichVlad/service-podof/pkg/pvz_grpc_pb/v1"
)

type GRPCServer struct {
    server   *grpc.Server
    listener net.Listener
    logger   *slog.Logger
	config *config.Config
}

func NewGRPCServer(
    service *service.Service,
    logger *slog.Logger,
	config *config.Config,
) (*GRPCServer, error) {
    listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GrpcPort))
    if err != nil {
        return nil, fmt.Errorf("failed to create listener: %w", err)
    }

    server := grpc.NewServer()
    

    pvzController := grpccontroller.NewPVZController(service, logger)
    pb.RegisterPVZServiceServer(server, pvzController)
    
  
    reflection.Register(server)

    return &GRPCServer{
        server:   server,
        listener: listener,
        logger:   logger,
    }, nil
}

func (s *GRPCServer) Start() error {
    s.logger.Info("starting gRPC server", slog.String("address", s.listener.Addr().String()))
    
    if err := s.server.Serve(s.listener); err != nil {
        return fmt.Errorf("failed to serve gRPC: %w", err)
    }
    
    return nil
}

func (s *GRPCServer) Stop() {
    s.logger.Info("stopping gRPC server")
    s.server.GracefulStop()
}