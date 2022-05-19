package zdpgo_datetime

import (
	"fmt"
	"testing"
)

/*
@Time : 2022/5/19 23:13
@Author : 张大鹏
@File : datetime_test
@Software: Goland2021.3.1
@Description: 日期时间相关测试
*/

func getDateTime() *DateTime {
	return New()
}

// 通用测试
func TestDateTime_common(t *testing.T) {
	d := getDateTime()

	// 今天
	fmt.Println(d.TimeNow)
	fmt.Println(d.TimeNow.ToString())                      // 今天此刻
	fmt.Println(d.TimeNow.ToDateTimeString())              // 今天此刻
	fmt.Println(d.TimeNow.ToDateString())                  // 今天日期
	fmt.Println(d.TimeNow.ToTimeString())                  // 今天时间
	fmt.Println(d.TimeNow.Now(NewYork).ToDateTimeString()) // 指定时区的今天此刻
	fmt.Println(d.TimeNow.Timestamp())                     // 今天秒级时间戳
	fmt.Println(d.TimeNow.TimestampMilli())                // 今天毫秒级时间戳
	fmt.Println(d.TimeNow.TimestampMicro())                // 今天微秒级时间戳
	fmt.Println(d.TimeNow.TimestampNano())                 // 今天纳秒级时间戳

	// 昨天
	fmt.Println(d.TimeYesterday)
	fmt.Println(d.TimeYesterday.ToString())
	fmt.Println(d.TimeYesterday.ToDateTimeString())
	fmt.Println(d.TimeYesterday.ToDateString())
	fmt.Println(d.TimeYesterday.ToTimeString())
	fmt.Println(d.TimeYesterday.Now(NewYork).ToDateTimeString())
	fmt.Println(d.TimeYesterday.Timestamp())
	fmt.Println(d.TimeYesterday.TimestampMilli())
	fmt.Println(d.TimeYesterday.TimestampMicro())
	fmt.Println(d.TimeYesterday.TimestampNano())

	// 明天
	fmt.Println(d.TimeTomorrow)
	fmt.Println(d.TimeTomorrow.ToString())
	fmt.Println(d.TimeTomorrow.ToDateTimeString())
	fmt.Println(d.TimeTomorrow.ToDateString())
	fmt.Println(d.TimeTomorrow.ToTimeString())
	fmt.Println(d.TimeTomorrow.Now(NewYork).ToDateTimeString())
	fmt.Println(d.TimeTomorrow.Timestamp())
	fmt.Println(d.TimeTomorrow.TimestampMilli())
	fmt.Println(d.TimeTomorrow.TimestampMicro())
	fmt.Println(d.TimeTomorrow.TimestampNano())
}
