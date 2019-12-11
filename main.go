package main 

import (
	"exclaim"
	"yell"
	"fmt"
	"token"
)


func main() {
    fmt.Println("Sanity Check")
    fmt.Println(yell.Yell("Sanity Check"))
    fmt.Println(exclaim.Exclaim("Sanity Check"))
    fmt.Println(token.Yell("Sanity Check"))
}
