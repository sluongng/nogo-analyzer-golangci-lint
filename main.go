package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/golangci/golangci-lint/pkg/golinters"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

var (
	template = `# Generated file, do not modify
package {{ .PackageName }}

import (
	"github.com/golangci/golangci-lint/pkg/golinters"
)

// Assuming that there is only 1 analyzer for each linter
var Analyzer = golinters.{{ .FuncName }}.GetAnalyzers()[0]`

	linters = []func() *goanalysis.Linter{
		golinters.NewAsciicheck,
		golinters.NewBodyclose,
		golinters.NewContainedCtx,
		golinters.NewContextCheck,
		golinters.NewDeadcode,
		golinters.NewDepguard,
		golinters.NewDogsled,
		golinters.NewDupl,
		golinters.NewDurationCheck,
		golinters.NewErrName,
		golinters.NewErrcheck,
		golinters.NewExportLoopRef,
		golinters.NewForbidigo,
		golinters.NewForceTypeAssert,
		golinters.NewFunlen,
		golinters.NewGoHeader,
		golinters.NewGoPrintfFuncName,
		golinters.NewGochecknoglobals,
		golinters.NewGochecknoinits,
		golinters.NewGocognit,
		golinters.NewGoconst,
		golinters.NewGocritic,
		golinters.NewGocyclo,
		golinters.NewGodot,
		golinters.NewGodox,
		golinters.NewGoerr113,
		golinters.NewGofmt,
		golinters.NewGofumpt,
		golinters.NewGoimports,
		golinters.NewGolint,
		golinters.NewGomodguard,
		golinters.NewIneffassign,
		golinters.NewInterfacer,
		golinters.NewLLL,
		golinters.NewMakezero,
		golinters.NewMaligned,
		golinters.NewMisspell,
		golinters.NewNakedret,
		golinters.NewNestif,
		golinters.NewNilErr,
		golinters.NewNoLintLint,
		golinters.NewNoctx,
		golinters.NewParallelTest,
		golinters.NewPrealloc,
		golinters.NewPromlinter,
		golinters.NewRowsErrCheck,
		golinters.NewSQLCloseCheck,
		golinters.NewScopelint,
		golinters.NewStructcheck,
		golinters.NewTparallel,
		golinters.NewTypecheck,
		golinters.NewUnconvert,
		golinters.NewUnparam,
		golinters.NewVarcheck,
		golinters.NewWSL,
		golinters.NewWastedAssign,
		golinters.NewWhitespace,
	}
)

func main() {
	switch os.Args[2] {
	case "fun":
		printFuncNames()
		break
	case "count":
		printCount()
		break
	case "generate":
	default:
		generate()
		break
	}
}

func generate() {
	// TODO
}

func printFuncNames() {
	packs, err := parser.ParseDir(token.NewFileSet(), "vendor/github.com/golangci/golangci-lint/pkg/golinters", nil, 0)
	if err != nil {
		log.Fatalf("Could not parse directory")
	}

	funcs := []string{}
	for _, pack := range packs {
		for _, file := range pack.Files {
			for _, d := range file.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					// Function should have a 'New' prefix
					// and take in zero param
					// and only have 1 output
					if strings.HasPrefix(fn.Name.Name, "New") &&
						len(fn.Type.Params.List) == 0 &&
						len(fn.Type.Results.List) == 1 {
						funcs = append(funcs, fn.Name.Name)
					}
				}
			}
		}
	}

	sort.Strings(funcs)

	for _, fun := range funcs {
		fmt.Printf(`golinters.%s,
`, fun)
	}
}

func printCount() {
	for _, fun := range linters {
		linter := fun()
		fmt.Printf("%s %d\n", linter.Name(), len(linter.GetAnalyzers()))
	}
}
