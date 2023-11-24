package manager

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/arybolovlev/orkestrator/api/proto/worker"
	"github.com/arybolovlev/orkestrator/api/structs"
)

type workerManager struct {
	worker.WorkerServer
}

func (m *workerManager) RegisterWorker(ctx context.Context, req *worker.RegisterWorkerRequest) (*worker.RegisterWorkerResponse, error) {
	log.Println("new worker registry received:", req.Name)

	w := structs.Worker{Name: req.Name}
	if _, ok := Workers[req.Name]; ok {
		log.Printf("worker %s already registered\n", w.Name)
		w.ID = Workers[req.Name].ID
	} else {
		w.ID = uuid.New().String()
		log.Printf("register a new worker %s: %s\n", w.Name, w.ID)
		log.Println("new worker was successfully registered:", w.ID, w.Name)
	}

	Workers[req.Name] = w
	log.Println("total workers registered:", len(Workers))

	return &worker.RegisterWorkerResponse{
		Id:   w.ID,
		Name: w.Name,
	}, nil
}
