package main

import (
    "fmt"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    "firecloud-slave/cpu/cpuproto" 
    "firecloud-slave/cpu/linux"
)

const (
    // Address gRPC服务地址
    Address = "127.0.0.1:50052"
)

func main() {
    // 连接
    conn, err := grpc.Dial(Address, grpc.WithInsecure())

    if err != nil {
        grpclog.Fatalln(err)
    }

    defer conn.Close()
	
	// 获取CPU信息
	cpuinfos,err := linux.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		grpclog.Fatal("stat read fail")
	}
	
	cpunum := cpuinfos.NumCPU()
	fmt.Println(cpunum)
	
    // 初始化客户端
    c := cpuproto.NewCpuInfoClient(conn)

    // 调用方法
    reqBody := new(cpuproto.SaveRequest)
    reqBody.Cpunum = uint32(cpunum)
    r, err := c.Save(context.Background(), reqBody)
    if err != nil {
        grpclog.Fatalln(err)
    }

    fmt.Println(r.Code)
}
