package manager

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/arybolovlev/orkestrator/api/proto/client"
	"github.com/arybolovlev/orkestrator/api/proto/worker"
	"github.com/arybolovlev/orkestrator/api/structs"
)

var (
	Jobs    = map[string]structs.Job{}
	Workers = map[string]structs.Worker{}
)

func Run(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("falied to listen: %v", err)
	}

	gs := grpc.NewServer()
	client.RegisterClientServer(gs, &clientManager{})
	worker.RegisterWorkerServer(gs, &workerManager{})

	log.Printf("server listening at %v", ln.Addr())

	if err := gs.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("tot ziens!")
}
