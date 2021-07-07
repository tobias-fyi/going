package main

import (
	"fmt"
	"os"
)

// func main() {
// 	_, file := path.Split("go-bootcamp/01-basics/01-path-separator.go")
// 	fmt.Println("file:", file)
// }

// func main() {
// 	speed, force := 100, 2.5
// 	momentum := float64(speed) * force
// 	fmt.Println(momentum)
// }

func main() {
	fmt.Println("Path:", os.Args[0])
	fmt.Println("One:", os.Args[1])
	fmt.Println("Last:", os.Args[len(os.Args)-1])
}
