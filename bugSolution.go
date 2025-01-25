```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 5) // Buffered channel to prevent blocking

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := 0; i < 5; i++ {
                        ch <- i
                }
                close(ch)
        }()

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := range ch {
                        fmt.Println(i)
                }
        }()

        wg.Wait()
}
```