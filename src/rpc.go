package mapReduce

import (
	"os"
	"strconv"
)

type RegisterArgs struct {
	ID string
}

type RegisterReply struct {
	ID     string
	Status bool
}

type PingArgs struct {
}

type PingReply struct {
}

// Add your RPC definitions here.

// Cook up a unique-ish UNIX-domain socket name
// in /var/tmp, for the coordinator.
// Can't use the current directory since
// Athena AFS doesn't support UNIX-domain sockets.
func coordinatorSock() string {
	s := "/var/tmp/5840-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
