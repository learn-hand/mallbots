package main

import (
	"fmt"
	"os"

	"github.com/learn-hand/mallbots/internal/monolith"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg monolith.AppConfig
	m := monolith.Monolith{cfg: cfg}
	if err = m.startupModules(); err != nil {
		return err
	}
	return nil
}
