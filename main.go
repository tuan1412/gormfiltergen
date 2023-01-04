package main

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
)

const cmdUsage = `
Usage: gormfiltergen 
Examples:
	gormfiltergen -src=model.go -dest=filter.go
`

func main() {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports}
	pkgs, err := packages.Load(cfg, "file=./test/model.go")

	if err != nil {
		log.Fatalf("load: %v\n", err)
	}

	pkg := pkgs[0]

	// TODO:
	// Step 1: Load Structs
	// Step 2: Load Tags, filter Tag "filter"
	// Step 3: Check Field Type, type Basic or type Named Time

	objNames := pkg.Types.Scope().Names()
	structs := make([]*types.Struct, 0, len(objNames))

	for _, name := range objNames {
		obj := pkg.Types.Scope().Lookup(name)
		structType, ok := obj.Type().Underlying().(*types.Struct)
		if ok {
			structs = append(structs, structType)
		}
	}
	fmt.Println(len(structs))
}
