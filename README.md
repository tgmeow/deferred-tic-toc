Deferred Tic Toc
================

Go code (golang) package that provides an easy to use timer for printing the execution time of a function.

Example
----------------
```go
package yourPackage

import (
    "github.com/tgmeow/deferred-tic-toc"    
    "time"
)

func runSomething() {
    defer deferredTicToc.ticToc("something")()
    time.Sleep(time.Second)
}

func main() {
    runSomething()
}
```