package main

import (
	"echo-boilerplate/conf"
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	if err := conf.InitConfig("conf/config.toml"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	e := echo.New()

	// e = Init(e)
	// if conf.Conf.LogLevel == "DEBUG" {
	// 	fmt.Println("DEBUG MODE")
	// 	e.Debug = true
	// }
	// e.HideBanner = false

	e.Start(conf.Conf.Server.Addr)
}
