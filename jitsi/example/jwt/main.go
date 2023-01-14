package main

import (
	"micro_services/msbase/jitsi"

	"github.com/labstack/echo/v4"
)

func main() {
	ec := echo.New()

	jitsi.InitWithConfigFile("config.yaml", ec)

	ec.Logger.Fatal(ec.Start(":4000"))
}
