package manager

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/arybolovlev/orkestrator/proto/client"
)

type manager struct {
	client.ClientServer
}

func (m *manager) CreateJob(ctx context.Context, job *client.CreateJobRequest) (*client.CreateJobResponse, error) {
	log.Println("New job request received:", job.Name)

	r := &client.CreateJobResponse{
		Id:   uuid.New().String(),
		Name: job.Name,
	}

	log.Println("New job was successfully registered:", r.Id, r.Name)

	return r, nil
}
