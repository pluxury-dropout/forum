package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("write smth")
		return
	}
	for _, i := range os.Args[1:] {
		fmt.Println(i)
	}
}
