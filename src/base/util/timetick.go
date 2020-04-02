package util

import (
	//	"runtime/debug"
	"time"
)

type Timer struct {
	now   int64
	split int64
}

func NewTimer(ms uint32) *Timer {
	t := &Timer{
		now:   time.Now().UnixNano(),
		split: int64(ms) * int64(time.Millisecond),
	}

	return t
}

func (this *Timer) Check() bool {

	n := time.Now().UnixNano()

	if n < this.now {
		return false
	}

	s := n - this.now
	if s < this.split {
		return false
	}

	this.now = n
	return true
}

type ITimeTick interface {
	TimeAction()
}

type TimeTick struct {
	Derived ITimeTick
}

func (this *TimeTick) IsFinal() bool {
	return false
}

func (this *TimeTick) mainLoop() {

	defer func() {
		if err := recover(); err != nil {
			//log.Println("[异常] ", err, "\n", string(debug.Stack()))
		}
	}()

	t := time.Duration(0)
	for !this.IsFinal() {
		s := 10*time.Millisecond - t
		if s <= 0 {
			s = 1 * time.Millisecond
		}

		time.Sleep(s)

		n1 := time.Now().UnixNano()
		if this.Derived != nil {
			this.Derived.TimeAction()
		}
		n2 := time.Now().UnixNano()
		t = time.Duration(n2 - n1)
		if t > 10*time.Millisecond {
			//log.Println("timetick:", (n2-n1)/int64(time.Millisecond))
		}
	}
}

func (this *TimeTick) Run() {
	go this.mainLoop()
}

type MyFunctionTime struct {
}

func NewMyFunctionTime(str string, t time.Duration) {
	s := time.Now().UnixNano() - int64(t)
	if s < int64(10*time.Millisecond) {
		return
	}

	//log.Println("耗时统计:", str, " ms:", s/1e6)
}
