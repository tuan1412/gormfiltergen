package main

import (
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

	//objLength := pkg.Types.Scope().NumChildren()
	//
	//for i := 0; i < objLength; i += 1 {
	//	obj := pkg.Types.Scope().Child(i)
	//	fmt.Println(obj)
	//}
}
