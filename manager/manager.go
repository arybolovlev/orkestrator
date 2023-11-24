package manager

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/arybolovlev/orkestrator/api/job"
	"github.com/arybolovlev/orkestrator/api/proto/client"
)

var (
	Jobs = map[string]job.Job{}
)

func Run(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("falied to listen: %v", err)
	}

	gs := grpc.NewServer()
	client.RegisterClientServer(gs, &manager{})

	log.Printf("server listening at %v", ln.Addr())

	if err := gs.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("tot ziens!")
}
