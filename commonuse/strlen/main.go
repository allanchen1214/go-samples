package main

import (
	"fmt"
)

func main() {
	str := "aancADBBNN123"
	fmt.Printf("len(%s): %d\n", str, len(str))

	str = "helloworld你好世界"
	fmt.Printf("len(%s): %d\n", str, len(str))

	data := []byte(str)
	fmt.Printf("len data: %d\n", len(data))

}
