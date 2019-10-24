Deferred Tic Toc
================

Go code (golang) package that provides an easy to use timer for printing the execution time of a function.

Example
----------------
```go
package yourPackage

import (
    "time"
    "github.com/tgmeow/deferred-tic-toc"
)

func runSomething() {
    defer ticToc("something")
    time.Sleep(time.Second)
}

func main() {
    runSomething()
}
```