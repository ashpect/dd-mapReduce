package mapReduce

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type workerData struct {
	ID            string
	status        int // 0 for idle, 1 for busy
	filesLocation string
}
type Coordinator struct {
	PORT          int
	mapWorkers    map[string]*workerData
	reduceWorkers map[string]*workerData
	idleWorkers   map[string]*workerData
}

// the RPC argument and reply types are defined in rpc.go.
func (c *Coordinator) Ping(args *PingArgs, reply *PingReply) error {
	fmt.Println("Ping")
	return nil
}

// Register worker with the corrdinator
func (c *Coordinator) RegisterWorker(args *RegisterArgs, reply *RegisterReply) error {
	fmt.Println("Registering_Worker")
	// Add the worker to the idle workers
	if _, exists := c.idleWorkers[args.ID]; exists {
		err := fmt.Errorf("worker ID is empty; cannot register worker")
		reply.ID = ""
		reply.Status = false
		return err
	}

	if args.ID == "" {
		err := fmt.Errorf("worker ID is empty; cannot register worker")
		fmt.Println(err)
		reply.ID = ""
		reply.Status = false
		return err
	}

	// Add to idle workers
	c.idleWorkers[args.ID] = &workerData{
		ID: args.ID,
	}

	reply.ID = args.ID
	reply.Status = true
	fmt.Printf("Worker with ID %s successfully registered\n", args.ID)

	return nil
}

// Start a thread that listens for RPCs from worker.go
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	// For socket based Connection
	// sockname := coordinatorSock()
	// os.Remove(sockname)
	// fmt.Println(sockname)
	// l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (c *Coordinator) Done() bool {

	// Should ping all the workers to check if they are done
	// If all the workers are done, then return true
	return false
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		PORT:          1234,
		mapWorkers:    make(map[string]*workerData),
		reduceWorkers: make(map[string]*workerData),
		idleWorkers:   make(map[string]*workerData),
	}

	filesNumber := len(files)
	// Start filesNumber go routines
	// for i := 0; i < filesNumber; i++ {
	// 	fmt.Println(files[i])
	// }
	fmt.Println(filesNumber)

	c.server()
	return &c
}
