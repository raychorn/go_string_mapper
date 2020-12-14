package main

import (
    "fmt"
    "./mapper"
)

func main() {
    s := mapper.NewSkipString(3, "Aspiration.com")
    mapper.MapString(&s)
    fmt.Println(s)
}
