package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 600; i++ {
		fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
		time.Sleep(1 * time.Second)
	}
}
