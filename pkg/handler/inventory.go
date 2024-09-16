package handler

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"

	pb "github.com/aiorch/inventory-manager/pkg/inventoryservice"
)

type InvHandler struct {
	pb.UnimplementedInventoryManagerServer
	Cw *CacheWrapper
}

func (s InvHandler) UnaryInventoryRequest(_ context.Context, _ *pb.InventoryRequest) (*pb.InventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryInventoryRequest not implemented")
}

func (s InvHandler) ServerStreamingInventory(_ *pb.InventoryRequest, _ grpc.ServerStreamingServer[pb.InventoryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingInventory not implemented")
}
