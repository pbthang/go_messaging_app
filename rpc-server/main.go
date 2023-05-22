package main

import (
	"log"
	"os"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	rpc "github.com/pbthang/go_messaging_app/rpc-server/kitex_gen/rpc/imservice"
)

func main() {
	var connectString string
	if os.Getenv("ENV") == "PROD" {
		connectString = "etcd:2379"
	} else {
		connectString = "127.0.0.1:2379"
	}
	log.Println("connectString: ", connectString)
	r, err := etcd.NewEtcdRegistry([]string{connectString}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
