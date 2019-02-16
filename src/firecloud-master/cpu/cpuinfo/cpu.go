package cpuinfo

import (
	"fmt"
	"reflect"
    "golang.org/x/net/context"
    "firecloud-master/cpu/cpuproto"
)

// 定义CpuInfo接口
type CpuInfo struct{}

//实现CpuInfo接口
func (c CpuInfo) Save(ctx context.Context, in *cpuproto.SaveRequest) (*cpuproto.SaveReply, error) {
    resp := new(cpuproto.SaveReply)
		for _,val := range in.Pro{
			tv := reflect.ValueOf(val).Elem()
			typeOfT := tv.Type()
			for i :=0; i < tv.NumField(); i++ {
				fieldName := typeOfT.Field(i).Name
				fieldValue := tv.FieldByName(fieldName)
				fmt.Println(fieldName,fieldValue)
			}
		}
		resp.Code = 200
    return resp, nil
}
