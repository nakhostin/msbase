package webapp

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

var (
	EP = echo.New()
)

type Server struct {
	ListenerPort string
}

func NewClient() *Server {
	return new(Server)
}

func (rcv *Server) Run() {
	EP.Start(fmt.Sprintf(":%s", rcv.ListenerPort))
}

func (rcv *Server) SetPort(p string) {
	rcv.ListenerPort = p
}
