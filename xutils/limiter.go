package xutils

//限制频率
//可以用Free中途取消限制
import (
	"sync"
	"time"
)

var limiter sync.Map

//地址不多可以这么干,多了要有回收机制,gocache
//返回true,表示可以通行
//返回false,限制通行
func Limit(addr string, sec int64) bool {
	value, ok := limiter.Load(addr)
	if !ok {
		limiter.Store(addr, time.Now().Unix())
		return true
	}
	if value.(int64) >= time.Now().Unix()-sec {
		return false
	} else {
		limiter.Store(addr, time.Now().Unix())
		return true
	}

}
func Free(addr string) {
	limiter.Delete(addr)
}
