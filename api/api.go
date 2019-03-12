package api

import (
	"flag"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port       *string
	debug      *bool
	echoServer *echo.Echo
)

func init() {
	port = flag.String("port", "9000", "port for the service HTTP")
	debug = flag.Bool("debug", false, "mod of the debug")
}

func Make() *echo.Echo {
	flag.Parse()

	echoServer = echo.New()

	// Esconde o cabeçalho do Echo
	echoServer.HideBanner = true

	echoServer.Use(middleware.Recover())

	return echoServer
}

func Run() {
	// For Heroku Work
	porta := os.Getenv("PORT")
	os.Setenv("debug", strconv.FormatBool(*debug))

	if porta == "" {
		porta = *port
	}

	if *debug {
		echoServer.Use(middleware.Logger())
	}

	echoServer.Use(middleware.CORS())
	echoServer.Logger.Fatal(echoServer.Start(":" + porta))
}
