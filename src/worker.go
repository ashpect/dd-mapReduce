package mapReduce

import (
	"fmt"
	"hash/fnv"
	"log"
	"net/rpc"
	"os"
	"time"

	"github.com/google/uuid"
)

type KeyValue struct {
	Key   string
	Value string
}

type worker struct {
	ID           string
	status       int
	fileLocation map[int]string
}

// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

// main/mrworker.go calls this function.
func Worker(mapf func(string, string) []KeyValue,
	reducef func(string, []string) string) {

	// Generate a unique ID for the worker
	newId := uuid.New().String()
	fmt.Printf("Worker ID: %v\n", newId)

	// Create a new worker, in memory
	worker := worker{
		ID:           newId,
		fileLocation: make(map[int]string),
	}

	// Register with the coordinator
	args := RegisterArgs{
		ID: worker.ID,
	}
	reply := RegisterReply{}

	ok := call("Coordinator.RegisterWorker", &args, &reply)
	if ok {
		if reply.Status {
			fmt.Printf("Successfully registered worker with ID: %v\n", reply)
		} else {
			fmt.Printf("Failed to register worker with ID: %v\n", reply)
		}
	} else {
		fmt.Printf("COORDINATOR DOWN!\n")
	}

	// Start pinging go routine at the side
	done := make(chan struct{})
	go func() {
		fmt.Printf("Starting regular pings to coordinator\n")
		for {
			select {
			case <-done:
				fmt.Println("Goroutine exiting.")
				return
			default:
				if err := pingCoordinator(); err != nil {
					fmt.Printf("Error: %s\n", err)
					// Cleanup if writing files or something
					os.Exit(1)
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	close(done)
	return
}

func pingCoordinator() error {
	args := PingArgs{}
	reply := PingReply{}

	ok := call("Coordinator.Ping", &args, &reply)
	if ok {
		return nil
	} else {
		err := fmt.Errorf("COORDINATOR DOWN!")
		return err
	}
}

// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
func call(rpcname string, args interface{}, reply interface{}) bool {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	// sockname := coordinatorSock()
	// c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
