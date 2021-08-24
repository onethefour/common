package xutils

import "encoding/json"

//用来调试打印结构体,map,interface
func String(d interface{}) string {
	str, _ := json.Marshal(d)
	return string(str)
}
