package workers

import (
	"fmt"
)

// WorkerB is another concrete implementation of IWorker.
type WorkerB struct {
}

// DoWork implements the DoWork method for WorkerB.
func (w WorkerB) DoWork() {
	fmt.Println("WorkerB is doing work!")
}
