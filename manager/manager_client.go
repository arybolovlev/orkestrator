package manager

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/arybolovlev/orkestrator/api/job"
	"github.com/arybolovlev/orkestrator/api/proto/client"
	"github.com/arybolovlev/orkestrator/api/task"
)

type manager struct {
	client.ClientServer
}

func (m *manager) RegisterJob(ctx context.Context, req *client.RegisterJobRequest) (*client.RegisterJobResponse, error) {
	log.Println("new job request received:", req.Name)

	j := job.Job{Name: req.Name}
	if _, ok := Jobs[req.Name]; ok {
		log.Printf("job %s already exists\n", j.Name)
		j.ID = Jobs[req.Name].ID
	} else {
		j.ID = uuid.New().String()
		log.Printf("register a new job %s: %s\n", j.Name, j.ID)
	}

	for _, t := range req.Task {
		j.Tasks = append(j.Tasks, task.Task{
			Name: t.Name, Image: task.Image{
				Name: t.Image.Name,
				Tag:  &t.Image.Tag,
			},
		})
	}

	log.Println("validate received job")
	if err := j.Validate(); err != nil {
		log.Printf("job %s is not valid: %s", j.Name, err)
		return nil, err
	}
	log.Printf("job %s is valid", j.Name)

	Jobs[j.Name] = j
	log.Println("new job was successfully registered:", j.ID, j.Name)
	log.Println("total jobs registered:", len(Jobs))

	return &client.RegisterJobResponse{
		Id:   j.ID,
		Name: j.Name,
	}, nil
}
