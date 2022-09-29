package main

import (
	"fmt"
	"time"

	"github.com/panjf2000/gnet"
)

type EventHandler struct {
}

func (e EventHandler) OnInitComplete(server gnet.Server) (action gnet.Action) {
	return gnet.None
}

func (e EventHandler) OnShutdown(server gnet.Server) {}

func (e EventHandler) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return nil, 0
}

func (e EventHandler) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	return gnet.None
}

func (e EventHandler) PreWrite(c gnet.Conn) {
}

func (e EventHandler) AfterWrite(c gnet.Conn, b []byte) {
}

func (e EventHandler) React(packet []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if err := c.AsyncWrite(packet); err != nil {
		fmt.Println("react send to failed, err: ", err)
	}
	fmt.Println(string(packet))
	if string(packet) == "quit" {
		return nil, gnet.Close
	}
	return nil, 0
}

func (e EventHandler) Tick() (delay time.Duration, action gnet.Action) {
	return 0, 0
}

func main() {
	addr := "tcp://localhost:8081"
	if err := gnet.Serve(EventHandler{}, addr, gnet.WithMulticore(true), gnet.WithReusePort(true)); err != nil {
		panic(err)
	}
}
