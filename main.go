package main

import (
	"fmt"

	"git.tecnotree.com/business-enablement/domv6/go-spring.git/registry"
)

// init function initializes the workerRegistry.
func init() {
	registry.AutoRegisterWorkers("git.tecnotree.com/business-enablement/domv6/go-spring.git/workers")
	//registry.RegisterWorker("WorkerA", reflect.TypeOf(workers.WorkerA{}))
	//registry.RegisterWorker("WorkerB", reflect.TypeOf(workers.WorkerB{}))
}

func main() {
	fmt.Println(registry.Get())
	// Create instances of workers using the registry
	workerA, errA := registry.CreateWorker("WorkerA")
	if errA != nil {
		panic(errA.Error())
	}
	workerA.DoWork()

	workerB, errB := registry.CreateWorker("WorkerB")
	if errB != nil {
		panic(errB.Error())
	}
	workerB.DoWork()
}
