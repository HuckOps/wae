package handler

import (
	"context"
	"wae/ci"
	"wae/pkg/logger"
	"wae/service"
)

// GRPC Handler

type CIServer struct {
	ci.UnimplementedCiRunnerServiceServer
}

func (s *CIServer) RegisterRunner(ctx context.Context, req *ci.RegisterRunnerReq) (*ci.RegisterRunnerResp, error) {
	_, err := service.RegisterRunner(ctx, req)
	if err != nil {
		return &ci.RegisterRunnerResp{Success: false}, err
	}
	return &ci.RegisterRunnerResp{Success: true}, nil
}

func (s *CIServer) DispatchTask(stream ci.CiRunnerService_DispatchTaskStreamServer) error {
	return nil
}

func (s *CIServer) ExecuteTaskStream(stream ci.CiRunnerService_ExecuteTaskStreamServer) error {
	return nil
}

func (s *CIServer) Heartbeat(ctx context.Context, req *ci.HeartbeatReq) (*ci.HeartbeatResp, error) {
	err := service.UpdateHeartbeat(ctx, req.GetUuid())
	if err != nil {
		logger.Logger.Sugar().Errorf("Heartbeat update failed, uuid: %s, err: %s", req.GetUuid(), err.Error())
		return &ci.HeartbeatResp{Success: false}, err
	}
	logger.Logger.Sugar().Infof("Heartbeat update success, uuid: %s", req.GetUuid())
	return &ci.HeartbeatResp{Success: true}, nil
}
