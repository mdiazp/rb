package main

import (
	"fmt"

	"github.com/mdiazp/rb/server/conf"
	dbhandlers "github.com/mdiazp/rb/server/db/handlers"
)

var (
	config *conf.Configuration
	db     dbhandlers.Handler
)

func init() {
	var e error
	configPath := "/home/kino/my_configs/rb"
	config, e = conf.LoadConfiguration(configPath, "dev")
	if e != nil {
		panic(fmt.Errorf("Fail loading configuration: %s", e.Error()))
	}
	db, e = dbhandlers.NewHandler(config)

	if e != nil {
		panic(fmt.Errorf("Fail at dbhandlers.NewHandler: %s", e.Error()))
	}
}
