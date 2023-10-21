package main

import (
	"flag"
	"fmt"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/config"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/handler"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/socket"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/message/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	hub := socket.NewHub()
	go hub.Run()

	ctx := svc.NewServiceContext(c, hub)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
