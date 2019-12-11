package main 

import (
	"exclaim"
	"yell"
	"fmt"
)


func main() {
    fmt.Println("Sanity Check")
    fmt.Println(yell.Yell("Sanity Check"))
    fmt.Println(exclaim.Exclaim("Sanity Check"))
}
