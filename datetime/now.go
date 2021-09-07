// Package datetime 提供日期时间相关的功能，并提供了一个秒级精度的本地时间及相关变量，当仅需要秒级精度的时间时，可以直接从这里获取而不用调用time.Now()
// 该文件内提供了与日期时间相关的函数的当前时间版本
package datetime

// LawAgeNow 返回周岁
// param: birthday: 生日，TimeFormatDate 后的日期，比如"2011-11-11"
// return: 如果输入格式不正确，返回 InvalidAge, 否则返回周岁
func LawAgeNow(birthday string) int {
	return LawAge(DateNow(), birthday)
}

// WeekFirstDateNow  返回本星期星期一的日期，形如 2006-01-01
func WeekFirstDateNow() string {
	return WeekFirstDate(Now())
}

// SecPastToday ：获取今天已经过去了多少秒 [0, 86399]
func SecPastToday() int {
	return SecPastOfDay(UnixTime())
}

// SecRestToday : 获取今天还剩余多少秒 [1, 86400]
func SecRestToday() int {
	return SecRestOfDay(UnixTime())
}

// FirstUnixTimeToday 返回当天第一秒的时间戳
func FirstUnixTimeToday() int64 {
	return FirstUnixTimeOfDay(UnixTime())
}

// LastUnixTimeToday 返回当天最后一秒的时间戳
func LastUnixTimeToday(unixTime int64) int64 {
	return LastUnixTimeOfDay(UnixTime())
}