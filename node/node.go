package node

type Node struct {
	Name            string
	IP              string
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	TaskCount       int
}
