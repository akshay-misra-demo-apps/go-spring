package main

import (
	"fmt"
	"reflect"

	"git.tecnotree.com/business-enablement/domv6/go-spring.git/registry"
	"git.tecnotree.com/business-enablement/domv6/go-spring.git/workers"
)

// init function initializes the workerRegistry.
func init() {
	registry.RegisterWorker("WorkerA", reflect.TypeOf(workers.WorkerA{}))
	registry.RegisterWorker("WorkerB", reflect.TypeOf(workers.WorkerB{}))
}

func main() {
	fmt.Println(registry.Get())
	// Create instances of workers using the registry
	workerA, errA := registry.CreateWorker("WorkerA")
	if errA == nil {
		workerA.DoWork()
	}

	workerB, errB := registry.CreateWorker("WorkerB")
	if errB == nil {
		workerB.DoWork()
	}
}
