package main

import (
	"fmt"
	"go/types"
	"os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
	_ "go.opentelemetry.io/otel/metric"
	"golang.org/x/tools/go/packages"
)

func main() {
	path, err := filepath.Abs(filepath.Join("..", "..", "wrappers.go"))
	if err != nil {
		fmt.Println("Error constructing path:", err)
		os.Exit(1)
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	f := jen.NewFile("mocks")

	cfg := &packages.Config{Mode: packages.NeedTypes}
	pkgs, err := packages.Load(cfg, "go.opentelemetry.io/otel/metric/embedded")
	if err != nil {
		fmt.Println("Error loading package:", err)
		os.Exit(1)
	}

	if len(pkgs) == 0 {
		fmt.Println("No packages found.")
		os.Exit(1)
	}

	pkg := pkgs[0]

	for _, name := range pkg.Types.Scope().Names() {
		obj := pkg.Types.Scope().Lookup(name)
		if obj == nil {
			continue
		}

		if _, ok := obj.Type().Underlying().(*types.Interface); ok {
			f.Type().Id(name+"Wrapper").Struct(
				jen.Qual("go.opentelemetry.io/otel/metric/embedded", name),
				jen.Id("*Mock"+name),
			).Line()
		}
	}

	err = f.Render(file)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}
