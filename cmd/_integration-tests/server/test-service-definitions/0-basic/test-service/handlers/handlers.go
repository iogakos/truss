package handlers

import (
	"context"

	pb "github.com/metaverse/truss/cmd/_integration-tests/server/test-service-definitions/0-basic"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.TESTServer {
	return testService{}
}

type testService struct{}

// GetBasic implements Service.
func (s testService) GetBasic(ctx context.Context, in *pb.BasicTypeRequest) (*pb.BasicTypeResponse, error) {
	var resp pb.BasicTypeResponse
	resp = pb.BasicTypeResponse{
		// A:
		// B:
		// C:
		// D:
		// E:
		// F:
		// G:
		// H:
		// I:
		// J:
		// K:
		// L:
		// M:
		// N:
	}
	return &resp, nil
}

// PostBasic implements Service.
func (s testService) PostBasic(ctx context.Context, in *pb.BasicTypeRequest) (*pb.BasicTypeRequest, error) {
	var resp pb.BasicTypeRequest
	resp = pb.BasicTypeRequest{
		// A:
		// B:
		// C:
		// D:
		// E:
		// F:
		// G:
		// H:
		// I:
		// J:
		// K:
		// L:
		// M:
		// N:
	}
	return &resp, nil
}
