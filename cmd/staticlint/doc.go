// staticlint implements multiple needed static checks.
//
// Following set includes:
//
// 1. All checks from golang.org/x/tools/go/analysis/passes
//
// 2. All SA checks from https://staticcheck.io/docs/checks/
//
// 3. Check bytes buffer conversions via https://staticcheck.io/docs/checks/#S1030
//
// 4. Check wrapping errors https://github.com/fatih/errwrap
//
// 5. Check for database query in loops https://github.com/masibw/goone
//
// 6. Check for calling os.Exit in main func of main package
//
// Example:
//
//	staticlint -SA1012 <project path>
//
// Perform SA1012 analysis for given project.
// For more details run:
//
//	staticlint -help
//
// osexitfrommain checks if in main package have ever been called os.Exit() in func main().To run this check use following command:
//
//	staticlint -osexitfrommain
package main
