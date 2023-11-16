package workers

import (
	"fmt"
)

// WorkerA is a concrete implementation of IWorker.
type WorkerA struct {
}

// DoWork implements the DoWork method for WorkerA.
func (w WorkerA) DoWork() {
	fmt.Println("WorkerA is doing work!")
}
