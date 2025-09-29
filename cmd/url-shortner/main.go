package main

import (
	"fmt"
	"url-shortner/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
