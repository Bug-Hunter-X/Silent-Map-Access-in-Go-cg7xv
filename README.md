# Silent Map Access in Go

This repository demonstrates a common, yet subtle, error in Go: accessing a non-existent key in a map.  Unlike some other languages that might throw an exception, Go silently returns the zero value for the type.

This can lead to bugs that are hard to track down, as there's no immediate indication that an error has occurred.

The `bug.go` file shows the problematic code, and `bugSolution.go` provides solutions to handle this scenario gracefully.