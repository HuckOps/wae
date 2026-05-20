package service

import (
	"context"
	"fmt"
	"time"
	"wae/ci"
	"wae/db"
	"wae/model"
	"wae/repo"
)

func RegisterRunner(ctx context.Context, req *ci.RegisterRunnerReq) (*model.Runner, error) {
	runners, err := repo.GetByOptions[model.Runner](ctx, repo.WithWhere("uuid = ?", req.GetUuid()))
	if err != nil {
		return nil, err
	}

	var runner *model.Runner
	if len(runners) == 0 {
		runner = &model.Runner{
			UUID:    req.GetUuid(),
			Version: req.GetVersion(),
			OS:      req.GetOs(),
			Arch:    req.GetArch(),
			CPU:     req.GetCpu(),
		}
		err = repo.Create(ctx, runner)
		if err != nil {
			return nil, err
		}
	} else {
		runner = &runners[0]
	}

	err = UpdateHeartbeat(ctx, req.GetUuid())
	if err != nil {
		return runner, err
	}

	return runner, nil
}

const HEARTBEAT_PREFIX = "runner:heartbeat:"
const HEARTBEAT_EXPIRE = time.Minute
const HEARTBEAT_TOLERANCE = time.Second * 30

func UpdateHeartbeat(ctx context.Context, uuid string) error {
	return db.RedisClient.Set(ctx, fmt.Sprintf("%s%s", HEARTBEAT_PREFIX, uuid), time.Now(), HEARTBEAT_EXPIRE+HEARTBEAT_TOLERANCE).Err()
}
