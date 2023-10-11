package main

import (
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/app/delivery"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}

func main() {
	delivery.Server().Run()
}
