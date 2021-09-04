package xutils

//对一个字符串加锁,
//例如地址不能并发:
//Lock(address)
//defer Unlock(address)

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	mapLockers map[string]*Mutex
	tmplock    sync.RWMutex
)

type Mutex struct {
	sync.Mutex
	LockWait int64
}

func init() {
	mapLockers = make(map[string]*Mutex)
}

//等待锁的任务太多就返回失败
func LockMax(key string, max int64) error {
	tmplock.Lock()
	v, ok := mapLockers[key]
	if ok && v.LockWait >= max {
		tmplock.Unlock()
		return errors.New(fmt.Sprintf("(key:%v,等待锁任务数量超过上限:%v)", key, max))
	}
	if !ok {
		fmt.Println(key)
		v = &Mutex{}
		mapLockers[key] = v
	}
	tmplock.Unlock()
	atomic.AddInt64(&v.LockWait, 1)
	v.Lock()
	return nil
}
func Lock(key string) {
	tmplock.Lock()
	v, ok := mapLockers[key]
	if !ok {
		v = &Mutex{}
		mapLockers[key] = v
	}
	tmplock.Unlock()
	atomic.AddInt64(&v.LockWait, 1)
	v.Lock()
}

func Unlock(key string) {
	tmplock.Lock()
	v, ok := mapLockers[key]
	tmplock.Unlock()

	if !ok {
		panic("锁不存在:" + key)
	}
	atomic.AddInt64(&v.LockWait, -1)
	v.Unlock()
}

//延迟一定时间释放锁
func UnlockDelay(key string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		Unlock(key)
	}()
}
