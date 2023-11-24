package client

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/arybolovlev/orkestrator/api/proto/client"
	"github.com/arybolovlev/orkestrator/api/structs"
)

type JobSpec struct {
	Job []structs.Job `hcl:"job,block"`
}

func Run(port int, spec string) {
	var jobs JobSpec
	err := DecodeFile(spec, &jobs)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := client.NewClientClient(conn)

	log.Println("Loading specification from file", spec)

	for _, job := range jobs.Job {
		req := &client.RegisterJobRequest{
			Name: job.Name,
		}
		for _, task := range job.Tasks {
			t := &client.Task{
				Name: task.Name,
				Image: &client.Image{
					Name: task.Image.Name,
				},
			}
			if task.Image.Tag == nil {
				t.Image.Tag = "latest"
			}
			req.Task = append(req.Task, t)
		}
		job, err := c.RegisterJob(context.Background(), req)
		if err != nil {
			log.Fatalf("falied to create job %s: %s", req.Name, err)
			os.Exit(1)
		}
		log.Println("new job sucessfully created:", job)
	}
}
