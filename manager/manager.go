package manager

import (
	"log"
	"net"

	"github.com/arybolovlev/orkestrator/proto/client"
	"google.golang.org/grpc"
)

func Run() {
	ln, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("Falied to listen: %v", err)
	}

	srv := &manager{}
	gs := grpc.NewServer()
	client.RegisterClientServer(gs, srv)

	log.Printf("server listening at %v", ln.Addr())

	if err := gs.Serve(ln); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("Tot ziens!")
}
