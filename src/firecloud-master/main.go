package main

import (
    "net"

     "firecloud-master/cpu/cpuproto" // 引入编译生成的包
     "firecloud-master/cpu/cpuinfo" // 引入cpu接口包
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

const (
    // Address gRPC服务地址
    Address = "127.0.0.1:50052"
)


func main() {
    var SaveCpuServ = cpuinfo.CpuInfo{}
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        grpclog.Fatalf("failed to listen: %v", err)
    }

    // 实例化grpc Server
    s := grpc.NewServer()

    // 注册HelloService
    cpuproto.RegisterCpuInfoServer(s, SaveCpuServ)

    fmt.Println("Listen on " + Address)

    s.Serve(listen)
}
