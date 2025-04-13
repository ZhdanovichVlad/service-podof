package grpc

import (
	"context"
	"log/slog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	pb "github.com/ZhdanovichVlad/service-podof/pkg/pvz_grpc_pb/v1"
)


type pvsService interface {
	GetAllPvz(ctx context.Context) ([]entity.Pvz, error)
}

type PVZController struct {
    pb.UnimplementedPVZServiceServer
    service pvsService
    logger  *slog.Logger
}

func NewPVZController(service pvsService, logger *slog.Logger) *PVZController {
    return &PVZController{
        service: service,
        logger:  logger,
    }
}

func (c *PVZController) GetPVZList(ctx context.Context, req *pb.GetPVZListRequest) (*pb.GetPVZListResponse, error) {
    c.logger.Info("handling GetPVZList request")

    
    pvzList, err := c.service.GetAllPvz(ctx)
    if err != nil {
        c.logger.Error("failed to get PVZ list", 
            slog.String("error", err.Error()),
        )
        return nil, status.Error(codes.Internal, "failed to get PVZ list")
    }

   
    response := &pb.GetPVZListResponse{
        Pvzs: make([]*pb.PVZ, 0, len(pvzList)),
    }

    for _, pvz := range pvzList {
        response.Pvzs = append(response.Pvzs, &pb.PVZ{
            Id:               pvz.Id.String(),
            City:            pvz.City,
            RegistrationDate: timestamppb.New(pvz.RegistrationDate),
        })
    }

    c.logger.Info("successfully got PVZ list", 
        slog.Int("count", len(response.Pvzs)),
    )

    return response, nil
}