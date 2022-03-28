package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("We are using Go %v running in %v! Sweet %v!\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v", text)
	fmt.Printf("We are using Go %v running in %v! Sweet %v!\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)

}
