package mr

//
// RPC definitions.
//

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

// worker to master
type register struct {
	host string
}

type registerReply struct {
	ok bool
}

type TaskFin struct {
	taskId string
}
type TaskFinReply struct {
	ok bool
}

// master to worker
type assignTask struct {
	taskId string
	taskType string
	path string
}

type assignTaskReply struct {
	ok bool
}