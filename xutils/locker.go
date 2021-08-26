package xutils

//对一个字符串加锁,
//例如地址不能并发:
//Lock(address)
//defer Unlock(address)

import (
	"sync"
)

var (
	mapLockers map[string]*sync.Mutex
	tmplock    sync.RWMutex
)

func init() {
	mapLockers = make(map[string]*sync.Mutex)
}

func Lock(key string) {
	tmplock.Lock()
	v, ok := mapLockers[key]
	if !ok {
		v = &sync.Mutex{}
		mapLockers[key] = v
	}
	tmplock.Unlock()
	v.Lock()
}

func Unlock(key string) {
	tmplock.Lock()
	v, ok := mapLockers[key]
	tmplock.Unlock()

	if !ok {
		panic("锁不存在:" + key)
	}
	v.Unlock()
}
