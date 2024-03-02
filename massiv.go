package main

import "fmt"

func reverseSlice(slice []interface{}) []interface{} {
    reversed := make([]interface{}, len(slice))
    for i := 0; i < len(slice); i++ {
        reversed[i] = slice[len(slice)-1-i]
    }
    return reversed
}

func main() {
    arr := []interface{}{"Hello", "how", "are", "you", "?"}
    reversed := reverseSlice(arr)
    fmt.Println(reversed) // Выведет ["?", "you", "are", "how", "Hello"]
}
