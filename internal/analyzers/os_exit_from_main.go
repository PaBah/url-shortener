package analyzers

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// OSExitFromMainAnalyzer - analyzer which check if os.Exit() call ever appeared in func main() in package main
var OSExitFromMainAnalyzer = &analysis.Analyzer{
	Name:     "osexitfrommain",
	Doc:      "check if os.Exit() call ever appeared in func main() in package main",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	//inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	//neededNodes := []ast.Node{
	//	(*ast.File)(nil),
	//	(*ast.FuncDecl)(nil),
	//	(*ast.SelectorExpr)(nil),
	//}
	//mainChecked := false
	//inspect.Preorder(neededNodes, func(n ast.Node) {
	//	switch x := n.(type) {
	//	case *ast.File:
	//		if x.Name.Name != "main" {
	//			return
	//		}
	//	case *ast.FuncDecl:
	//		isFuncMain := x.Name.Name == "main"
	//		if mainChecked {
	//			mainChecked = false
	//			return
	//		}
	//		mainChecked = isFuncMain
	//	case *ast.SelectorExpr:
	//		ident, ok := x.X.(*ast.Ident)
	//		if mainChecked && ok && ident.Name == "os" && x.Sel.Name == "Exit" {
	//			pass.Reportf(ident.NamePos, "os.Exit called in main func in main package")
	//			return
	//		}
	//	}
	//})
	for _, file := range pass.Files {
		packageIsMain, funcIsMain := false, false
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.File:
				packageIsMain = x.Name.Name == "main"
			case *ast.FuncDecl:
				funcIsMain = x.Name.Name == "main"
			case *ast.SelectorExpr:
				ident, ok := x.X.(*ast.Ident)
				if packageIsMain && funcIsMain && ok && ident.Name == "os" && x.Sel.Name == "Exit" {
					pass.Reportf(ident.NamePos, "os.Exit called in main func in main package")
				}
			}
			return true
		})
	}
	return nil, nil
}
