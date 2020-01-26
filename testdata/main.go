package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("TEST_ENV_1"))
	fmt.Println(os.Getenv("TEST_ENV_2"))
}
