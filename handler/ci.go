package handler

import (
	"context"
	"fmt"
	"wae/ci"
)

// GRPC Handler

type CIServer struct {
	ci.UnimplementedCiRunnerServiceServer
}

func (s *CIServer) RegisterRunner(ctx context.Context, req *ci.RegisterRunnerReq) (*ci.RegisterRunnerResp, error) {
	fmt.Println("RegisterRunner", req)
	runnerID := req.RunnerId
	fmt.Println("runnerID", runnerID)

	return &ci.RegisterRunnerResp{Success: true}, nil
}

func (s *CIServer) DispatchTask(stream ci.CiRunnerService_DispatchTaskStreamServer) error {
	return nil
}

func (s *CIServer) ExecuteTaskStream(stream ci.CiRunnerService_ExecuteTaskStreamServer) error {
	return nil
}

func (s *CIServer) Heartbeat(ctx context.Context, req *ci.HeartbeatReq) (*ci.HeartbeatResp, error) {
	return nil, nil
}
