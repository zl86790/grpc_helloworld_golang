package main

import (
	"helloworld/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()                                        // 创建grpc服务
	services.RegisterHiServiceServer(rpcServer, new(services.HiService)) // 注册，new方法中的是services包中定义的结构体HiService
	lis, _ := net.Listen("tcp", ":30010")                                // 开启一个监听端口
	rpcServer.Serve(lis)                                                 // 启动服务
}
