package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"wae/ci"
	pb "wae/ci"
	"wae/config"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := config.LoadCIConfig("/home/huck/wae/ci_config.yaml"); err != nil {
		log.Fatalf("Can't read runner config file. pls check.")
	}

	// Check storage dir
	if _, err := os.Stat(config.CIConf.StorageDir); err != nil {
		log.Fatalf("Storage dir %s not found. pls check.", config.CIConf.StorageDir)
	}

	// get or generate runner id
	runnerUUIDFile := config.CIConf.StorageDir + "/runner_uuid"

	runnerUUID := ""
	runnerUUIDBytes, err := os.ReadFile(runnerUUIDFile)
	if err != nil {
		log.Printf("Not found runner uuid file %s. will generate one.", runnerUUIDFile)
		uuid := uuid.New().String()
		runnerUUID = uuid
		err = os.WriteFile(runnerUUIDFile, []byte(runnerUUID), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		runnerUUID = string(runnerUUIDBytes)
	}
	config.RunnerUUID = runnerUUID

	conn, err := grpc.Dial(
		config.CIConf.Server,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewCiRunnerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RegisterRunner(ctx, client); err != nil {
		panic(err)
	}

	log.Printf("RegisterRunner success. runnerUUID: %s", runnerUUID)

	// Start heartbeat goroutine
	go StartHeartbeatLoop(client)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func StartHeartbeatLoop(client pb.CiRunnerServiceClient) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			if err := Heartbeat(ctx, client); err != nil {
				log.Printf("Heartbeat failed: %v", err)
			} else {
				log.Printf("Heartbeat success. runnerUUID: %s", config.RunnerUUID)
			}
			cancel()
		}
	}
}

func RegisterRunner(ctx context.Context, client pb.CiRunnerServiceClient) error {
	currentOS := runtime.GOOS
	arch := runtime.GOARCH
	cpu := runtime.NumCPU()
	req := &ci.RegisterRunnerReq{
		Uuid: config.RunnerUUID,
		Os:   currentOS,
		Arch: arch,
		Cpu:  int32(cpu),
	}
	resp, err := client.RegisterRunner(ctx, req)
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("RegisterRunner failed. resp: %v", resp)
	}
	return nil
}

func Heartbeat(ctx context.Context, client pb.CiRunnerServiceClient) error {
	req := &ci.HeartbeatReq{
		Uuid: config.RunnerUUID,
	}
	resp, err := client.Heartbeat(ctx, req)
	if err != nil {
		return err
	}
	if !resp.Success {
		return fmt.Errorf("Heartbeat failed. resp: %v", resp)
	}
	return nil
}
