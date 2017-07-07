package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_DateOf(t *testing.T) {
	timestamp := int64(1136185384)
	expected := "2006-01-02 15:03:04 +0800 CST"
	assert.Equal(t, expected, DateOf(timestamp))
}

func wrongDuration(beginTime time.Time, Duration, checkCycle time.Duration) bool {
	return time.Since(beginTime) < Duration ||
		Duration+checkCycle < time.Since(beginTime)
}

func Test_WaitFunc(t *testing.T) {
	beginTime := time.Now()
	checkCycle := time.Millisecond * 100

	waitCh, wait := WaitFunc(checkCycle, "Test_WaitFunc")

	wait()
	t.Log(time.Now())
	waitDuration := checkCycle
	if wrongDuration(beginTime, waitDuration, checkCycle) {
		t.Error("wait()在最小时间前结束了。")
	}

	updateCycle2 := time.Millisecond * 500
	waitCh <- updateCycle2
	wait()
	t.Log(time.Now())
	waitDuration += updateCycle2
	if wrongDuration(beginTime, waitDuration, checkCycle) {
		t.Error("wait()没能在updateCycle结束前，修改为更大的updateCycle")
	}

	updateCycle3 := time.Millisecond * 200
	go func() {
		time.Sleep(updateCycle3 / 2)
		waitCh <- updateCycle3
	}()
	wait()
	t.Log(time.Now())
	waitDuration += updateCycle3
	if wrongDuration(beginTime, waitDuration, checkCycle) {
		t.Error("wait()没能在updateCycle结束前，修改为更小的updateCycle")
	}
}
func Test_SleepFunc(t *testing.T) {
	beginTime := time.Now()
	waitDuration := 100 * time.Millisecond
	deltaDuration := 2 * time.Millisecond

	sleep := SleepFunc(waitDuration)
	sleep()
	assert.True(t, !wrongDuration(beginTime, waitDuration, deltaDuration),
		"sleep() 没有休眠正确的时间：%s", waitDuration)
}

func Test_ParseLocalTime(t *testing.T) {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05")
	nowPIT, _ := ParseLocalTime(nowStr)

	if now.Unix()-nowPIT.Unix() != 0 {
		t.Errorf("无法把%s转换成%s，而是转换成了%s", nowStr, now, nowPIT)
	}
}
