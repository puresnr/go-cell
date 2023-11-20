// Package datetime 提供日期时间相关的功能
package datetime

import (
	"fmt"
	"github.com/puresnr/go-cell/cerror"
	"sync"
	"time"
)

// nowLocal 是一个本地时区的秒级精度的时间同步变量，当仅需要秒级精度的时间时，可以从这里获取而不用调用time.Now()
// 注意：nowLocal 不适用于定时器任务中，因为无法保证时间更新和定时器执行的顺序
var localZoneTime *ZoneTime

func LocalZoneTime() *ZoneTime { return localZoneTime }

func init() {
	var err error
	localZoneTime, err = New("Local")
	if err != nil {
		panic(err.Error())
	}
}

// ZoneTime 提供一个秒级精度的时间同步变量，当仅需要秒级精度的时间时，可以从这里获取而不用调用time.Now()
// 注意：ZoneTime 不适用于定时器任务中，因为无法保证时间更新和定时器执行的顺序
type ZoneTime struct {
	// 一旦设置就不会改变
	zone   string
	offset int
	loc    *time.Location

	// 每秒都变化的
	now         time.Time
	nowTs       uint32 // 虽然说无论什么时区, unixtimestamp都是一样的, 但是为了时间戳和ZoneTime里的其它变量一致, 还是每个ZoneTime实例分别维护自己的时间戳吧
	nowDatetime string
	nowSec      int

	// 每分钟变化的
	nowMinute int

	// 每小时变化的
	nowHour int

	// 每天变化的
	nowDay            int
	nowDayOfUnixEpoch int // 获取今天是自1970-1-1以来的第几天
	nowWeekday        time.Weekday
	nowWeekdayISO     int // ISO weekday，星期日设为7，范围[1, 7]

	// 每月变化的
	nowMonth time.Month

	// 每年变化的
	nowYear int

	rwLock sync.RWMutex
}

func New(location string) (*ZoneTime, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return nil, cerror.Wrap(err)
	}

	now := time.Now().In(loc)

	zt := &ZoneTime{
		loc:         loc,
		now:         time.Now().In(loc),
		nowTs:       uint32(now.Unix()),
		nowDatetime: now.Format(TimeFormatDatetime),
		nowSec:      now.Second(),
		nowMinute:   now.Minute(),
		nowHour:     now.Hour(),
		nowDay:      now.Day(),
		nowWeekday:  now.Weekday(),
		nowMonth:    now.Month(),
		nowYear:     now.Year(),
		rwLock:      sync.RWMutex{},
	}

	zt.zone, zt.offset = now.Zone()
	zt.nowDayOfUnixEpoch = zt.DayOfUnixEpoch(zt.nowTs)
	zt.nowWeekdayISO = ToWeekdayISO(zt.nowWeekday)

	go func() {
		// 此时now不会发生变化，所以不加锁了
		time.Sleep(time.Second - time.Duration(now.Nanosecond()))

		zt.setData()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				zt.setData()
			}
		}
	}()

	return zt, nil
}

func (z *ZoneTime) setData() {
	z.rwLock.Lock()
	defer z.rwLock.Unlock()

	z.now = time.Now().In(z.loc)
	z.nowTs = uint32(z.now.Unix())
	z.nowSec = z.now.Second()
	old := z.nowDatetime
	z.nowDatetime = z.now.Format(TimeFormatDatetime)

	if old[0:4] != z.nowDatetime[0:4] {
		z.onYearChange()
		return
	}

	if old[5:7] != z.nowDatetime[5:7] {
		z.onMonthChange()
		return
	}

	if old[8:10] != z.nowDatetime[8:10] {
		z.onDayChange()
		return
	}

	if old[11:13] != z.nowDatetime[11:13] {
		z.onHourChange()
		return
	}

	if old[14:16] != z.nowDatetime[14:16] {
		z.onMinuteChange()
		return
	}
}

func (z *ZoneTime) onYearChange() {
	z.nowYear = z.now.Year()

	z.onMonthChange()
}

func (z *ZoneTime) onMonthChange() {
	z.nowMonth = z.now.Month()

	z.onDayChange()
}

func (z *ZoneTime) onDayChange() {
	z.nowDay = z.now.Day()
	z.nowDayOfUnixEpoch = z.DayOfUnixEpoch(z.nowTs)
	z.nowWeekday = z.now.Weekday()
	z.nowWeekdayISO = ToWeekdayISO(z.nowWeekday)

	z.onHourChange()
}

func (z *ZoneTime) onHourChange() {
	z.nowHour = z.now.Hour()

	z.onMinuteChange()
}

func (z *ZoneTime) onMinuteChange() {
	z.nowMinute = z.now.Minute()
}

func (z *ZoneTime) Loc() *time.Location { return z.loc }

func (z *ZoneTime) Zone() string {
	return z.zone
}

func (z *ZoneTime) Offset() int {
	return z.offset
}

func (z *ZoneTime) Now() time.Time {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.now
}

func (z *ZoneTime) NowTs() uint32 {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowTs
}

func (z *ZoneTime) NowHour() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowHour
}

func (z *ZoneTime) NowMinute() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowMinute
}

func (z *ZoneTime) NowSec() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowSec
}

func (z *ZoneTime) NowDay() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowDay
}

func (z *ZoneTime) NowMonth() time.Month {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowMonth
}

func (z *ZoneTime) NowYear() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowYear
}

func (z *ZoneTime) NowDate() string {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowDatetime[0:10]
}

func (z *ZoneTime) NowDatetime() string {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowDatetime
}

func (z *ZoneTime) NowDayOfUnixEpoch() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowDayOfUnixEpoch
}

func (z *ZoneTime) NowWeekday() time.Weekday {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowWeekday
}

func (z *ZoneTime) NowWeekdayISO() int {
	z.rwLock.RLock()
	defer z.rwLock.RUnlock()

	return z.nowWeekdayISO
}

// NowLawAge 返回当前时间的周岁
// param: birthday: 生日，TimeFormatDate 后的日期，比如"2011-11-11"
// return: 如果输入格式不正确，返回 InvalidAge, 否则返回周岁
func (z *ZoneTime) NowLawAge(birthday string) int {
	return LawAge(z.NowDate(), birthday)
}

// NowWeekFirstDate  返回本星期星期一的日期，形如 2006-01-01
func (z *ZoneTime) NowWeekFirstDate() string {
	return WeekFirstDate(z.Now())
}

// SecPastToday ：获取今天已经过去了多少秒 [0, 86399]
func (z *ZoneTime) SecPastToday() uint32 {
	return z.SecPastOfDay(z.NowTs())
}

// SecRestToday : 获取今天还剩余多少秒 [1, 86400]
func (z *ZoneTime) SecRestToday() uint32 {
	return z.SecRestOfDay(z.NowTs())
}

// FirstDayListOfMonth ：返回当前时区一个时间段内的所有月份第一天的列表(包括 start 和 end 所在月份)，形如2019-03-01
func (z *ZoneTime) FirstDayListOfMonth(sYear, sMonth, eYear, eMonth int) (ret []string) {
	sFirst, _ := time.ParseInLocation(TimeFormatDate, fmt.Sprintf("%4d-%02d-%02d", sYear, sMonth, 1), z.loc)
	eFirst, _ := time.ParseInLocation(TimeFormatDate, fmt.Sprintf("%4d-%02d-%02d", eYear, eMonth, 1), z.loc)

	for first := sFirst; !first.After(eFirst); first = first.AddDate(0, 1, 0) {
		ret = append(ret, first.Format(TimeFormatDate))
	}

	return
}

// SecPastOfDay ：获取指定的时间戳在其所在的那一天已经过去了多少秒 [0, 86399]
func (z *ZoneTime) SecPastOfDay(ts uint32) uint32 {
	offset := int(ts) + z.Offset()
	if offset < 0 {
		return uint32(SecOneDay + offset)
	} else {
		return uint32(offset % SecOneDay)
	}
}

// SecRestOfDay : 获取指定的时间戳在其所在的那一天还剩余多少秒 [1, 86400]
func (z *ZoneTime) SecRestOfDay(ts uint32) uint32 {
	return SecOneDay - z.SecPastOfDay(ts)
}

// DayOfUnixEpoch : 获取指定时间戳是自1970-1-1以来的第几天，当时间戳在1970-1-1之前会返回什么，尚未测试
func (z *ZoneTime) DayOfUnixEpoch(ts uint32) int {
	return (int(ts) + z.Offset()) / SecOneDay
}
