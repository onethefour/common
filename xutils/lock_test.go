package xutils

import (
	"testing"
	"time"
)

func Test_lock(t *testing.T) {
	Lock("aaaaa")
	go func() {
		if err := LockMax("aaaaa", 2); err != nil {
			t.Log("1" + err.Error())
		} else {
			t.Log("获得锁1")
		}
	}()
	go func() {
		if err := LockMax("aaaaa", 2); err != nil {
			t.Log("2" + err.Error())
		} else {
			t.Log("获得锁2")
		}
	}()
	go func() {
		if err := LockMax("aaaaa", 2); err != nil {
			t.Log("3" + err.Error())
		} else {
			t.Log("获得锁3")
		}
	}()
	time.Sleep(time.Second)
	t.Log("准备解锁")
	UnlockDelay("aaaaa", time.Second)
	t.Log("解锁")
	time.Sleep(3 * time.Second)
}
