// Package datetime 提供日期时间相关的功能，并提供了一个秒级精度的本地时间及相关变量，当仅需要秒级精度的时间时，可以直接从这里获取而不用调用time.Now()
// 该文件维护了一个本地时间，每秒更新一次，同时更新相关的一些常用变量
package datetime

import (
	"sync"
	"time"
)

var (
	// 这两个值一旦设置就不会改变
	zone string
	offsetSec int

	// 每秒都变化的
	now time.Time
	unixTime int64
	datetimeNow string
	secondNow int

	// 每分钟变化的
	minuteNow int

	// 每小时变化的
	hourNow int

	// 每天变化的
	dayNow int
	dayOfUnixEpochNow int     // 获取今天是自1970-1-1以来的第几天
	weekday time.Weekday
	weekdayISO int    // ISO weekday，星期日设为7，范围[1, 7]

	// 每月变化的
	monthNow time.Month

	// 每年变化的
	yearNow int

	rwLock sync.RWMutex
)

func init() {
	initData()

	go func() {
		// 此时now不会发生变化，所以不加锁了
		time.Sleep(time.Second - time.Duration(now.Nanosecond()))

		setData()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <- ticker.C:
				setData()
			}
		}
	}()
}

func initData() {
	// 初始化时不会并发

	zone, offsetSec = now.Zone()

	now = time.Now()
	unixTime = now.Unix()
	secondNow = now.Second()
	minuteNow = now.Minute()
	hourNow = now.Hour()
	yearNow, monthNow, dayNow = now.Date()
	datetimeNow = now.Format(TimeFormatDatetime)

	dayOfUnixEpochNow = DayOfUnixEpoch(unixTime)

	weekday = now.Weekday()
	weekdayISO = ToWeekdayISO(weekday)
}

func onYearChange() {
	yearNow = now.Year()

	onMonthChange()
}

func onMonthChange() {
	monthNow = now.Month()

	onDayChange()
}

func onDayChange() {
	dayNow = now.Day()
	dayOfUnixEpochNow = DayOfUnixEpoch(unixTime)
	weekday = now.Weekday()
	weekdayISO = ToWeekdayISO(weekday)

	onHourChange()
}

func onHourChange() {
	hourNow = now.Hour()

	onMinuteChange()
}

func onMinuteChange() {
	minuteNow = now.Minute()
}

func setData() {
	rwLock.Lock()
	defer rwLock.Unlock()

	now = time.Now()
	unixTime = now.Unix()
	secondNow = now.Second()
	old := datetimeNow
	datetimeNow = now.Format(TimeFormatDatetime)

	if old[0:4] != datetimeNow[0:4] {
		onYearChange()
		return
	}

	if old[5:7] != datetimeNow[5:7] {
		onMonthChange()
		return
	}

	if old[8:10] != datetimeNow[8:10] {
		onDayChange()
		return
	}

	if old[11:13] != datetimeNow[11:13] {
		onHourChange()
		return
	}

	if old[14:16] != datetimeNow[14:16] {
		onMinuteChange()
		return
	}
}

func Zone() string {
	return zone
}

func OffsetSec() int {
	return offsetSec
}

func Now() time.Time {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return now
}

func UnixTime() int64 {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return unixTime
}

func HourNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return hourNow
}

func MinuteNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return minuteNow
}

func SecondNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return secondNow
}

func DayNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return dayNow
}

func MonthNow() time.Month {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return monthNow
}

func YearNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return yearNow
}

func DateNow() string {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return datetimeNow[0:10]
}

func DatetimeNow() string {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return datetimeNow
}

func DayOfUnixEpochNow() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return dayOfUnixEpochNow
}

func Weekday() time.Weekday {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return weekday
}

func WeekdayISO() int {
	rwLock.RLock()
	defer rwLock.RUnlock()

	return weekdayISO
}