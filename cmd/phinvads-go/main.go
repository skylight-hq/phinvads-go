package main

import (
	"github.com/skylight-hq/phinvads-go/internal/app"
	"github.com/skylight-hq/phinvads-go/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	pv := app.SetupApp(cfg)

	pv.Run()
}
