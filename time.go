package util

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// DateOf 返回一个unix tiemstamp的格式化。
func DateOf(t int64) string {
	return time.Unix(t, 0).String()
}

// WaitFunc 返回 一个 cycleCh  chan<- time.Duration 和 一个 Wait 函数
// 如果 cycleCh <- d ，则 Wait() 的sleep时间被修改为 d
// sleep时间的修改会立即生效，不用等到此waitDuration结束。
// 例如：当100秒的waitDuration时间已过51秒时，把waitDuration修改为50秒，程序会立刻结束。
//      不用等到100秒才结束。
// 例如：100秒的waitDuration在结束前，把waitDuration修改为200秒，
//      此次Wait()会sleep 200秒才结束。
//
// checkCycle是检查是否到期的时间段，也是最小等待时间段。
//
// 当checkCycle较大时，Wait()的sleep时间会和waitDuration，有较大的差距，
// 程序的实际等待时间，总是 checkCycle × int(waitDuration/checkCycle + 1)
func WaitFunc(checkCycle time.Duration, name string) (chan<- time.Duration, func()) {
	cycleCh := make(chan time.Duration, 7)
	beginTime := time.Now()
	waitDuration := checkCycle

	return cycleCh, func() {

		for {
			select {
			case waitDuration = <-cycleCh:
				if waitDuration <= checkCycle {
					log.Printf("WARNING: %s的 等待时间 <= 检查周期 ，程序只会按照检查周期来等待。", name)
				}
				log.Printf("INFO: %s的 等待时间 已经修改为：%s", name, waitDuration)
			default:
			}

			// 把判断是否结束的语句，放在最后，很有必要。
			// 因为很有可能，wait()的调用周期总是大于waitDuration
			// 而导致总是无法进入for循环，来改变waitDuration的值
			// 特别是第一个waitDuration的值为checkCycle，总是比较小的。

			if time.Now().Before(beginTime.Add(waitDuration)) {
				time.Sleep(checkCycle)
			} else {
				break
			}
		}

		beginTime = time.Now()
	}
}

//SleepFunc 返回一个等待sleep函数， 使程序暂停一个duration。
//SleepFunc是以上WaitFunc的简化版本，通常运用于API访问限制
//利用闭包，把beginTime变量包裹到了sleep函数内。
func SleepFunc(duration time.Duration) func() {
	beginTime := time.Now()
	return func() {
		//无需判断，如果Sleep的时间为负，则不会Sleep。
		time.Sleep(duration - time.Since(beginTime))
		beginTime = time.Now()
	}
}

// ParseLocalTime 把string格式的时间，转换成time.time
// NOTICE: 我写这个函数的原因是
// 1. 添加更详解的错误说明。
// 2. 编写单元测试，查看是否真的转换成功了。
func ParseLocalTime(strTime string) (*time.Time, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", strTime, time.Local)
	if err != nil {
		msg := fmt.Sprintf("无法把%s转换成time.time格式: %s", strTime, err)
		return nil, errors.New(msg)
	}

	return &t, nil
}
