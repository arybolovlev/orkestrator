package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/arybolovlev/orkestrator/task"
	"github.com/docker/docker/client"
)

func main() {
	// t := task.Task{
	// 	ID:            uuid.New(),
	// 	Name:          "task",
	// 	State:         task.Pending,
	// 	Image:         "orkestrator",
	// 	RestartPolicy: task.Always,
	// 	Memory:        1024,
	// 	Disk:          1,
	// }
	// te := task.TaskEvent{
	// 	ID:        uuid.New(),
	// 	State:     task.Pending,
	// 	Timestamp: time.Now(),
	// 	Task:      t,
	// }
	// fmt.Printf("task: %v\n", t)
	// fmt.Printf("task event: %v\n", te)

	// w := worker.Worker{
	// 	Queue: *queue.New(),
	// 	DB:    make(map[uuid.UUID]task.Task),
	// }
	// fmt.Printf("worker: %v\n", w)
	// w.CollectStats()
	// w.RunTask()
	// w.StartTask()
	// w.StopTask()

	// m := manager.Manager{
	// 	Pending: *queue.New(),
	// 	TaskDB:  make(map[string][]task.Task),
	// 	EventDB: make(map[string][]task.TaskEvent),
	// 	Workers: []string{
	// 		w.Name,
	// 	},
	// }
	// fmt.Printf("manager: %v\n", m)
	// m.SelectWorker()
	// m.UpdateTasks()
	// m.SendWork()

	// n := node.Node{
	// 	Name:   "node",
	// 	IP:     "192.168.1.1",
	// 	Memory: 4096,
	// 	Disk:   10,
	// 	Role:   node.Worker,
	// }
	// fmt.Printf("node: %v\n", n)

	c := task.Config{
		Name: "orkestrator-container-a",
		Image: task.Image{
			Repository: "postgres",
			Tag:        "15.1",
		},
		Env: []string{
			"POSTGRES_USER=cube",
			"POSTGRES_PASSWORD=secret",
		},
	}

	dc, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Printf("Failed to create a Docker client: %v\n", err)
	}
	d := task.Docker{
		Client: dc,
		Config: c,
	}
	result := d.Run()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
		os.Exit(1)
	}
	fmt.Printf("Container %s is running with config %v\n", result.ContainerID, c)

	time.Sleep(5 * time.Second)

	result = d.Stop()
	if result.Error != nil {
		fmt.Printf("%v\n", result.Error)
	}
	fmt.Printf("Container %s has been stopped and removed\n", result.ContainerID)
}
