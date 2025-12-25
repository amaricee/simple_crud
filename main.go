package main

import (
	"fmt"
	"github.com/amaricee/simple_crud/config"

)

func main() {
	fmt.Println("Starting application...")

	config.ConnectDB()
}