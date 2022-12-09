package task

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type Config struct {
	Name          string
	Image         Image
	Memory        int64
	Disk          int64
	Env           []string
	RestartPolicy RestartPolicy
}

type Docker struct {
	Client      *client.Client
	Config      Config
	ContainerID string
}

type DockerResult struct {
	Error       error
	Action      string
	ContainerID string
	Result      string
}

type Image struct {
	Repository string
	Tag        string
}

type RestartPolicy string

const (
	No            RestartPolicy = "no"
	Always        RestartPolicy = "always"
	UnlessStopped RestartPolicy = "unless-stopped"
	OnFailure     RestartPolicy = "on-failure"
)

type State int

const (
	Pending State = iota
	Scheduled
	Completed
	Running
	Failed
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy RestartPolicy
	StartTime     time.Time
	FinishTime    time.Time
}

type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}

func imageName(i Image) string {
	return fmt.Sprintf("%s:%s", i.Repository, i.Tag)
}

func (d *Docker) Run() DockerResult {
	ctx := context.Background()
	image := imageName(d.Config.Image)
	reader, err := d.Client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		log.Printf("Error pulling image %s: %v\n", image, err)
		return DockerResult{Error: err}
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)

	resp, err := d.Client.ContainerCreate(ctx, &container.Config{
		Image: image,
		Env:   d.Config.Env,
	}, &container.HostConfig{
		RestartPolicy: container.RestartPolicy{
			Name: string(d.Config.RestartPolicy),
		},
		Resources: container.Resources{
			Memory: d.Config.Memory,
		},
		PublishAllPorts: true,
	}, nil, nil, d.Config.Name)
	if err != nil {
		log.Printf("Error creating container using image %s: %v\n", image, err)
		return DockerResult{Error: err}
	}

	d.ContainerID = resp.ID

	err = d.Client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("Error starting container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}

	out, err := d.Client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		log.Printf("Error getting logs for container %s: %v\n", resp.ID, err)
		return DockerResult{Error: err}
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return DockerResult{
		ContainerID: resp.ID,
		Action:      "start",
		Result:      "success",
		Error:       nil,
	}
}

func (d *Docker) Stop() DockerResult {
	ctx := context.Background()
	log.Printf("Attempting to stop container %v", d.ContainerID)
	err := d.Client.ContainerStop(ctx, d.ContainerID, nil)
	if err != nil {
		log.Printf("Error stopping container %s: %v\n", d.ContainerID, err)
		return DockerResult{Error: err}
	}

	ro := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         false,
	}

	err = d.Client.ContainerRemove(ctx, d.ContainerID, ro)
	if err != nil {
		log.Printf("Error removing container %s: %v\n", d.ContainerID, err)
		return DockerResult{Error: err}
	}

	return DockerResult{
		ContainerID: d.ContainerID,
		Action:      "stop",
		Result:      "success",
		Error:       nil,
	}
}
