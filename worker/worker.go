package worker

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arybolovlev/orkestrator/api/proto/worker"
)

func Run(name string, port int) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	w := worker.NewWorkerClient(conn)

	wrk, err := w.RegisterWorker(context.Background(), &worker.RegisterWorkerRequest{Name: name})
	if err != nil {
		log.Fatalf("falied to subscribe worker %s: %s", name, err)
		os.Exit(1)
	}
	log.Println("new worker sucessfully subscribed:", wrk)
}
