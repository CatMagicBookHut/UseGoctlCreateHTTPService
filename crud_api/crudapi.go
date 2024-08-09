package main

import (
	"flag"
	"fmt"

	"crud_api/internal/config"
	"crud_api/internal/handler"
	"crud_api/internal/svc"
	"crud_api/model"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/crudapi-api.yaml", "the config file")

func main() {
	// 连接数据库
	model.InitDb()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
