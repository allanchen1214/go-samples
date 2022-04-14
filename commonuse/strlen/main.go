package main

import (
	"fmt"
)

func main() {
	str := "aancADBBNN123"
	fmt.Printf("len(%s): %d\n", str, len(str)) // 13

	str = "helloworld你好世界"
	fmt.Printf("len(%s): %d\n", str, len(str)) // 22

	data := []byte(str)
	fmt.Printf("len data: %d\n", len(data)) // 22

}
