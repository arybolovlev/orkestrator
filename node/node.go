package node

type NodeRole string

const (
	Worker NodeRole = "worker"
)

type Node struct {
	Name            string
	IP              string
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	TaskCount       int
	Role            NodeRole
}
