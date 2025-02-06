```go
package main

import "fmt"

func main() {
    var m map[string]int
    value, ok := m["foo"] // Use the 'ok' idiom
    if ok {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found")
    }

    //Alternatively, use a conditional assignment
    value = m["foo"]
    if value == 0 { // 0 is the zero value for int
        fmt.Println("Key likely not found. Value is 0")
    }
}
```