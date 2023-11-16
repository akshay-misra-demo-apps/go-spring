package registry

import (
	"fmt"
	"reflect"

	"git.tecnotree.com/business-enablement/domv6/go-spring.git/api"
)

// workerRegistry is a map to store worker types.
var workerRegistry = make(map[string]reflect.Type)

func Get() map[string]reflect.Type {
	return workerRegistry
}

// RegisterWorker registers a worker type with the registry.
func RegisterWorker(name string, workerType reflect.Type) {
	workerRegistry[name] = workerType
}

// CreateWorker creates an instance of a registered worker by name.
func CreateWorker(name string) (api.IWorker, error) {
	if workerType, ok := workerRegistry[name]; ok {
		// Create an instance of the registered worker type using reflection
		workerValue := reflect.New(workerType).Elem()
		// Check if the created instance satisfies the IWorker interface
		if worker, ok := workerValue.Interface().(api.IWorker); ok {
			return worker, nil
		}
	}
	return nil, fmt.Errorf("worker %s is not registered", name)
}
