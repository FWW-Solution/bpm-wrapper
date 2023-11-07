package main

import (
	"bpm-wrapper/internal/container"

	"bpm-wrapper/internal/config"
)

func main() {
	cfg := config.InitConfig()
	container.InitService(cfg)
}
