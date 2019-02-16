package main

import (
    "fmt"
    "reflect"
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

    var processor_dict  = new(cpuproto.Processor)
    var processor_list []*cpuproto.Processor

  	// 获取CPU信息
  	cpuinfos,err := linux.ReadCPUInfo("/proc/cpuinfo")
  	if err != nil {
  		grpclog.Fatal("stat read fail")
  	}

    // 遍历
    // s := reflect.ValueOf(processor_dict).Elem()
    // typeOfT := s.Type()
    // for i := 0; i < s.NumField(); i++ {
    //     // f := s.Field(i)
    //     // fmt.Printf("%d %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
    //     fmt.Printf("%s\n",typeOfT.Field(i).Name)
    // }

    // 遍历cpuinfos.Processors[i] 即 Processor
    t1t := reflect.TypeOf(cpuinfos.Processors[0])
    t1v := reflect.ValueOf(cpuinfos.Processors[0])
    //processor_dict是一个指针，reflect.ValueOf(processor_dict)，reflect.ValueOf(*processor_dict) 传入的都只是拷贝
    //Elem()表示传入指针引用，其它非指针需要加取地址符，表示传入引用 reflect.ValueOf(&cpuinfos.Processors[0])
    t2v := reflect.ValueOf(processor_dict).Elem()
    for j :=0; j < t1t.NumField(); j++ {
      fieldName := t1t.Field(j).Name
      fieldValue := t1v.FieldByName(fieldName)
      t2v.FieldByName(fieldName).Set(fieldValue)
    }
    processor_dict.Id = 1
    processor_list = append(processor_list,processor_dict)
    fmt.Println(cpuinfos.Processors[0].Id)

    // 初始化客户端
    c := cpuproto.NewCpuInfoClient(conn)

    // 调用方法
    reqBody := new(cpuproto.SaveRequest)
    reqBody.Ip = "192.168.10.3"
    reqBody.Pro = processor_list
    r, err := c.Save(context.Background(), reqBody)
    if err != nil {
        grpclog.Fatalln(err)
    }

    fmt.Println(r.Code)
}
