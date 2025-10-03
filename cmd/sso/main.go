package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()

	fmt.Println(cfg)
}
