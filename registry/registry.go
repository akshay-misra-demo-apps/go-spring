package registry

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"unsafe"

	"golang.org/x/tools/go/packages"

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

// AutoRegisterWorkers dynamically registers all types implementing IWorker from the specified package.
func AutoRegisterWorkers(pkgPath string) error {
	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, pkgPath)
	if err != nil {
		return fmt.Errorf("failed to load packages: %v", err)
	}

	fmt.Println("Found packages: ", pkgs)

	for _, pkg := range pkgs {
		for _, syntax := range pkg.Syntax {
			for _, decl := range syntax.Decls {
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for _, spec := range genDecl.Specs {
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {
							obj := pkg.TypesInfo.ObjectOf(typeSpec.Name)
							if typ, ok := obj.Type().(*types.Named); ok {
								RegisterWorker(typeSpec.Name.Name, reflect.TypeOf((*types.Named)(unsafe.Pointer(typ))).Elem())
							}
						}
					}
				}
			}
		}
	}

	return nil
}

// CreateWorker creates an instance of a registered worker by name.
func CreateWorker(name string) (api.IWorker, error) {
	fmt.Println("CreateWorker:: ", workerRegistry)
	if workerType, ok := workerRegistry[name]; ok {
		// Create an instance of the registered worker type using reflection
		workerValue := reflect.New(workerType).Elem()
		// Check if the created instance satisfies the IWorker interface
		if worker, ok := workerValue.Interface().(api.IWorker); ok {
			return worker, nil
		} else {
			fmt.Println("worker is not of type 'api.IWorker'.")
		}
	}
	return nil, fmt.Errorf("worker %s is not registered", name)
}
