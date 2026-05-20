package main

import (
	"context"
	"log"
	"os"
	"os/signal"
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
	runnerIDFile := config.CIConf.StorageDir + "/runner_id"

	runnerID := ""
	runnerIDBytes, err := os.ReadFile(runnerIDFile)
	if err != nil {
		log.Printf("Not found runner id file %s. will generate one.", runnerIDFile)
		uuid := uuid.New().String()
		runnerID = uuid
		err = os.WriteFile(runnerIDFile, []byte(runnerID), 0644)
		if err != nil {
			panic(err)
		}
	} else {
		runnerID = string(runnerIDBytes)
	}
	config.RunnerID = runnerID

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

	resp, err := client.RegisterRunner(ctx, &ci.RegisterRunnerReq{
		RunnerId: runnerID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("RegisterRunner success: %v. runnerID: %s", resp.Success, runnerID)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
