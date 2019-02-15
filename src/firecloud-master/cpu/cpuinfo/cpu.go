package cpuinfo

import (
	"fmt"
    "golang.org/x/net/context"
    "firecloud-master/cpu/cpuproto"
)

// 定义CpuInfo接口
type CpuInfo struct{}

//实现CpuInfo接口
func (c CpuInfo) Save(ctx context.Context, in *cpuproto.SaveRequest) (*cpuproto.SaveReply, error) {
    resp := new(cpuproto.SaveReply)
    fmt.Println(in.Cpunum)
    return resp, nil
}

