package handler

import (
	context "context"

	"github.com/redis/go-redis"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	"aiorch/inventory-manager/pkg/inventoryservice"
)

type invHandler struct {
	cw *CacheWrapper
}

func (s *invHandler) UnaryInventoryRequest(_ context.Context, *pb.InventoryRequest) (*pb.InventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryInventoryRequest not implemented")
}

func (s *invHandler) ServerStreamingInventory(req *pb.InventoryRequest, s grpc.ServerStreamingServer[InventoryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingInventory not implemented")
}