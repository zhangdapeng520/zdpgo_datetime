package main

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"log"
	"time"
)

/*
@Time : 2022/5/19 22:48
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description:
*/

func main() {
	// 从秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestamp(-1).ToString()         // 1970-01-01 07:59:59 +0800 CST
	carbon.CreateFromTimestamp(0).ToString()          // 1970-01-01 08:00:00 +0800 CST
	carbon.CreateFromTimestamp(1).ToString()          // 1970-01-01 08:00:01 +0800 CST
	carbon.CreateFromTimestamp(1649735755).ToString() // 2022-04-12 11:55:55 +0800 CST

	// 从毫秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestampMilli(1649735755981).ToString() // 2022-04-12 11:55:55.981 +0800 CST
	// 从微秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestampMicro(1649735755981566).ToString() // 2022-04-12 11:55:55.981566 +0800 CST
	// 从纳秒级时间戳创建 Carbon 实例
	carbon.CreateFromTimestampNano(1649735755981566000).ToString() // 2022-04-12 11:55:55.981566 +0800 CST

	// 从年月日时分秒创建 Carbon 实例
	carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
	// 从年月日时分秒创建 Carbon 实例，包含毫秒
	carbon.CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
	// 从年月日时分秒创建 Carbon 实例，包含微秒
	carbon.CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
	// 从年月日时分秒创建 Carbon 实例，包含纳秒
	carbon.CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)
	carbon.CreateFromDate(2020, 8, 5).ToString() // // 2020-08-05 13:14:15 +0800 CST
	// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)，包含毫秒
	carbon.CreateFromDateMilli(2020, 8, 5, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
	// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)，包含微秒
	carbon.CreateFromDateMicro(2020, 8, 5, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
	// 从年月日创建 Carbon 实例(时分秒默认为当前时分秒)，包含纳秒
	carbon.CreateFromDateNano(2020, 8, 5, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)
	carbon.CreateFromTime(13, 14, 15).ToString() // 2020-08-05 13:14:15 +0800 CST
	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)，包含毫秒
	carbon.CreateFromTimeMilli(13, 14, 15, 999).ToString() // 2020-08-05 13:14:15.999 +0800 CST
	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)，包含微秒
	carbon.CreateFromTimeMicro(13, 14, 15, 999999).ToString() // 2020-08-05 13:14:15.999999 +0800 CST
	// 从时分秒创建 Carbon 实例(年月日默认为当前年月日)，包含纳秒
	carbon.CreateFromTimeNano(13, 14, 15, 999999999).ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("").ToDateTimeString()                    // 空字符串
	carbon.Parse("0").ToDateTimeString()                   // 空字符串
	carbon.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空字符串
	carbon.Parse("0000-00-00").ToDateTimeString()          // 空字符串
	carbon.Parse("00:00:00").ToDateTimeString()            // 空字符串

	carbon.Parse("2020-08-05").ToString()                    // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("2020-08-05 13:14:15").ToString()           // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999").ToString()       // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999999").ToString()    // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("2020-08-05 13:14:15.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("2020-08-05T13:14:15+08:00").ToString()           // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999+08:00").ToString()       // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999+08:00").ToString()    // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.Parse("20200805").ToString()                 // 2020-08-05 00:00:00 +0800 CST
	carbon.Parse("20200805131415").ToString()           // 2020-08-05 13:14:15 +0800 CST
	carbon.Parse("20200805131415.999").ToString()       // 2020-08-05 13:14:15.999 +0800 CST
	carbon.Parse("20200805131415.999999").ToString()    // 2020-08-05 13:14:15.999999 +0800 CST
	carbon.Parse("20200805131415.999999999").ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	carbon.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString()                     // 2020-08-05 13:14:15
	carbon.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
	carbon.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString()           // 2020-08-05 13:14:15
	carbon.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", carbon.Tokyo).ToDateTimeString()       // 2020-08-05 14:14:15

	carbon.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString()               // 2020-08-05 13:14:15
	carbon.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString()   // 2020-08-05 13:14:15
	carbon.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString()     // 2020-08-05 13:14:15
	carbon.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

	// 将 time.Time 转换成 Carbon
	carbon.Time2Carbon(time.Now())
	// 将 Carbon 转换成 time.Time
	carbon.Now().Carbon2Time()

	// 本世纪开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
	// 本世纪结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

	// 本年代开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
	carbon.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
	carbon.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
	// 本年代结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
	carbon.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
	carbon.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

	// 本年开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
	// 本年结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

	// 本季度开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
	// 本季度结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

	// 本月开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
	// 本月结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

	// 本周开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString()                                // 2020-08-02 00:00:00
	carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
	carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
	// 本周结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString()                                // 2020-08-08 23:59:59
	carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
	carbon.Parse("2020-08-05 13:14:15").SetWeekStartsAt(carbon.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

	// 本日开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
	// 本日结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

	// 本小时开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
	// 本小时结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

	// 本分钟开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
	// 本分钟结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

	// 本秒开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfSecond().ToString() // 2020-08-05 13:14:15 +0800 CST
	// 本秒结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfSecond().ToString() // 2020-08-05 13:14:15.999999999 +0800 CST

	// 三个世纪后
	carbon.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
	// 三个世纪后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
	// 一个世纪后
	carbon.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
	// 一个世纪后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
	// 三个世纪前
	carbon.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
	// 三个世纪前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
	// 一个世纪前
	carbon.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
	// 一世纪前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

	// 三个年代后
	carbon.Parse("2020-02-29 13:14:15").AddDecades(3).ToDateTimeString() // 2050-03-01 13:14:15
	// 三个年代后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
	// 一个年代后
	carbon.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
	// 一个年代后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
	// 三个年代前
	carbon.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
	// 三个年代前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
	// 一个年代前
	carbon.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
	// 一个年代前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

	// 三年后
	carbon.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
	// 三年后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
	// 一年后
	carbon.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
	// 一年后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
	// 三年前
	carbon.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
	// 三年前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
	// 一年前
	carbon.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
	// 一年前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

	// 三个季度后
	carbon.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
	// 三个季度后(月份不溢出)
	carbon.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15
	// 一个季度后
	carbon.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
	// 一个季度后(月份不溢出)
	carbon.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
	// 三个季度前
	carbon.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
	// 三个季度前(月份不溢出)
	carbon.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
	// 一个季度前
	carbon.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
	// 一个季度前(月份不溢出)
	carbon.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

	// 三个月后
	carbon.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
	// 三个月后(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
	// 一个月后
	carbon.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
	// 一个月后(月份不溢出)
	carbon.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
	// 三个月前
	carbon.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
	// 三个月前(月份不溢出)
	carbon.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
	// 一个月前
	carbon.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
	// 一个月前(月份不溢出)
	carbon.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

	// 三周后
	carbon.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
	// 一周后
	carbon.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
	// 三周前
	carbon.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
	// 一周前
	carbon.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

	// 三天后
	carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
	// 一天后
	carbon.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
	// 三天前
	carbon.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
	// 一天前
	carbon.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

	// 三小时后
	carbon.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
	// 二小时半后
	carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString()  // 2020-08-05 15:44:15
	carbon.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
	// 一小时后
	carbon.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
	// 三小时前
	carbon.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
	// 二小时半前
	carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString()  // 2020-08-05 10:44:15
	carbon.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
	// 一小时前
	carbon.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

	// 三分钟后
	carbon.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
	// 二分钟半后
	carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString()  // 2020-08-05 13:16:45
	carbon.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
	// 一分钟后
	carbon.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
	// 三分钟前
	carbon.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
	// 二分钟半前
	carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString()  // 2020-08-05 13:11:45
	carbon.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
	// 一分钟前
	carbon.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

	// 三秒钟后
	carbon.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
	// 二秒钟半后
	carbon.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
	// 一秒钟后
	carbon.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
	// 三秒钟前
	carbon.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
	// 二秒钟半前
	carbon.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
	// 一秒钟前
	carbon.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14

	// 三毫秒后
	carbon.Parse("2020-08-05 13:14:15.222222222").AddMilliseconds(3).ToString() // 2020-08-05 13:14:15.225222222 +0800 CST
	// 一毫秒后
	carbon.Parse("2020-08-05 13:14:15.222222222").AddMillisecond().ToString() // 2020-08-05 13:14:15.223222222 +0800 CST
	// 三毫秒前
	carbon.Parse("2020-08-05 13:14:15.222222222").SubMilliseconds(3).ToString() // 2020-08-05 13:14:15.219222222 +0800 CST
	// 一毫秒前
	carbon.Parse("2020-08-05 13:14:15.222222222").SubMillisecond().ToString() // 2020-08-05 13:14:15.221222222 +0800 CST

	// 三微秒后
	carbon.Parse("2020-08-05 13:14:15.222222222").AddMicroseconds(3).ToString() // 2020-08-05 13:14:15.222225222 +0800 CST
	// 一微秒后
	carbon.Parse("2020-08-05 13:14:15.222222222").AddMicrosecond().ToString() // 2020-08-05 13:14:15.222223222 +0800 CST
	// 三微秒前
	carbon.Parse("2020-08-05 13:14:15.222222222").SubMicroseconds(3).ToString() // 2020-08-05 13:14:15.222219222 +0800 CST
	// 一微秒前
	carbon.Parse("2020-08-05 13:14:15.222222222").SubMicrosecond().ToString() // 2020-08-05 13:14:15.222221222 +0800 CST

	// 三纳秒后
	carbon.Parse("2020-08-05 13:14:15.222222222").AddNanoseconds(3).ToString() // 2020-08-05 13:14:15.222222225 +0800 CST

	// 相差多少年
	carbon.Parse("2021-08-05 13:14:15").DiffInYears(carbon.Parse("2020-08-05 13:14:15")) // -1
	// 相差多少年（绝对值）
	carbon.Parse("2021-08-05 13:14:15").DiffAbsInYears(carbon.Parse("2020-08-05 13:14:15")) // 1

	// 相差多少月
	carbon.Parse("2020-08-05 13:14:15").DiffInMonths(carbon.Parse("2020-07-05 13:14:15")) // -1
	// 相差多少月（绝对值）
	carbon.Parse("2020-08-05 13:14:15").DiffAbsInMonths(carbon.Parse("2020-07-05 13:14:15")) // 1

	// 相差多少周
	carbon.Parse("2020-08-05 13:14:15").DiffInWeeks(carbon.Parse("2020-07-28 13:14:15")) // -1
	// 相差多少周（绝对值）
	carbon.Parse("2020-08-05 13:14:15").DiffAbsInWeeks(carbon.Parse("2020-07-28 13:14:15")) // 1

	// 相差多少天
	carbon.Parse("2020-08-05 13:14:15").DiffInDays(carbon.Parse("2020-08-04 13:14:15")) // -1
	// 相差多少天（绝对值）
	carbon.Parse("2020-08-05 13:14:15").DiffAbsInDays(carbon.Parse("2020-08-04 13:14:15")) // 1

	// 相差多少小时
	carbon.Parse("2020-08-05 13:14:15").DiffInHours(carbon.Parse("2020-08-05 12:14:15")) // -1
	// 相差多少小时（绝对值）
	carbon.Parse("2020-08-05 13:14:15").DiffAbsInHours(carbon.Parse("2020-08-05 12:14:15")) // 1

	// 相差多少分
	carbon.Parse("2020-08-05 13:14:15").DiffInMinutes(carbon.Parse("2020-08-05 13:13:15")) // -1
	// 相差多少分（绝对值）
	carbon.Parse("2020-08-05 13:14:15").DiffAbsInMinutes(carbon.Parse("2020-08-05 13:13:15")) // 1

	// 相差多少秒
	carbon.Parse("2020-08-05 13:14:15").DiffInSeconds(carbon.Parse("2020-08-05 13:14:14")) // -1
	// 相差多少秒（绝对值）
	carbon.Parse("2020-08-05 13:14:15").DiffAbsInSeconds(carbon.Parse("2020-08-05 13:14:14")) // 1

	// 相差字符串
	carbon.Now().DiffInString()                       // just now
	carbon.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
	carbon.Now().SubYearsNoOverflow(1).DiffInString() // 1 year
	// 相差字符串（绝对值）
	carbon.Now().DiffAbsInString(carbon.Now())                       // just now
	carbon.Now().AddYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year
	carbon.Now().SubYearsNoOverflow(1).DiffAbsInString(carbon.Now()) // 1 year

	// 对人类友好的可读格式时间差
	carbon.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
	carbon.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
	carbon.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
	carbon.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
	carbon.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

	carbon.Parse("2020-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year before
	carbon.Parse("2019-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years before
	carbon.Parse("2018-08-05 13:14:15").DiffForHumans(carbon.Now()) // 1 year after
	carbon.Parse("2022-08-05 13:14:15").DiffForHumans(carbon.Now()) // 2 years after

	// 是否是零值时间
	carbon.Parse("").IsZero()                              // true
	carbon.Parse("0").IsZero()                             // true
	carbon.Parse("0000-00-00 00:00:00").IsZero()           // true
	carbon.Parse("0000-00-00").IsZero()                    // true
	carbon.Parse("00:00:00").IsZero()                      // true
	carbon.Parse("2020-08-05 00:00:00").IsZero()           // false
	carbon.Parse("2020-08-05").IsZero()                    // false
	carbon.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

	// 是否是有效时间
	carbon.Parse("").IsValid()                              // false
	carbon.Parse("0").IsValid()                             // false
	carbon.Parse("0000-00-00 00:00:00").IsValid()           // false
	carbon.Parse("0000-00-00").IsValid()                    // false
	carbon.Parse("00:00:00").IsValid()                      // false
	carbon.Parse("2020-08-05 00:00:00").IsValid()           // true
	carbon.Parse("2020-08-05").IsValid()                    // true
	carbon.Parse("2020-08-05").SetTimezone("xxx").IsValid() // false

	// 是否是无效时间
	carbon.Parse("").IsInvalid()                              // true
	carbon.Parse("0").IsInvalid()                             // true
	carbon.Parse("0000-00-00 00:00:00").IsInvalid()           // true
	carbon.Parse("0000-00-00").IsInvalid()                    // true
	carbon.Parse("00:00:00").IsInvalid()                      // true
	carbon.Parse("2020-08-05 00:00:00").IsInvalid()           // false
	carbon.Parse("2020-08-05").IsInvalid()                    // false
	carbon.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

	// 是否是当前时间
	carbon.Now().IsNow() // true
	// 是否是未来时间
	carbon.Tomorrow().IsFuture() // true
	// 是否是过去时间
	carbon.Yesterday().IsPast() // true

	// 是否是闰年
	carbon.Parse("2020-08-05 13:14:15").IsLeapYear() // true
	// 是否是长年
	carbon.Parse("2020-08-05 13:14:15").IsLongYear() // true

	// 是否是一月
	carbon.Parse("2020-08-05 13:14:15").IsJanuary() // false
	// 是否是二月
	carbon.Parse("2020-08-05 13:14:15").IsFebruary() // false
	// 是否是三月
	carbon.Parse("2020-08-05 13:14:15").IsMarch() // false
	// 是否是四月
	carbon.Parse("2020-08-05 13:14:15").IsApril() // false
	// 是否是五月
	carbon.Parse("2020-08-05 13:14:15").IsMay() // false
	// 是否是六月
	carbon.Parse("2020-08-05 13:14:15").IsJune() // false
	// 是否是七月
	carbon.Parse("2020-08-05 13:14:15").IsJuly() // false
	// 是否是八月
	carbon.Parse("2020-08-05 13:14:15").IsAugust() // false
	// 是否是九月
	carbon.Parse("2020-08-05 13:14:15").IsSeptember() // true
	// 是否是十月
	carbon.Parse("2020-08-05 13:14:15").IsOctober() // false
	// 是否是十一月
	carbon.Parse("2020-08-05 13:14:15").IsNovember() // false
	// 是否是十二月
	carbon.Parse("2020-08-05 13:14:15").IsDecember() // false

	// 是否是周一
	carbon.Parse("2020-08-05 13:14:15").IsMonday() // false
	// 是否是周二
	carbon.Parse("2020-08-05 13:14:15").IsTuesday() // true
	// 是否是周三
	carbon.Parse("2020-08-05 13:14:15").IsWednesday() // false
	// 是否是周四
	carbon.Parse("2020-08-05 13:14:15").IsThursday() // false
	// 是否是周五
	carbon.Parse("2020-08-05 13:14:15").IsFriday() // false
	// 是否是周六
	carbon.Parse("2020-08-05 13:14:15").IsSaturday() // false
	// 是否是周日
	carbon.Parse("2020-08-05 13:14:15").IsSunday() // false

	// 是否是工作日
	carbon.Parse("2020-08-05 13:14:15").IsWeekday() // false
	// 是否是周末
	carbon.Parse("2020-08-05 13:14:15").IsWeekend() // true

	// 是否是昨天
	carbon.Parse("2020-08-04 13:14:15").IsYesterday() // true
	carbon.Parse("2020-08-04 00:00:00").IsYesterday() // true
	carbon.Parse("2020-08-04").IsYesterday()          // true
	// 是否是今天
	carbon.Parse("2020-08-05 13:14:15").IsToday() // true
	carbon.Parse("2020-08-05 00:00:00").IsToday() // true
	carbon.Parse("2020-08-05").IsToday()          // true
	// 是否是明天
	carbon.Parse("2020-08-06 13:14:15").IsTomorrow() // true
	carbon.Parse("2020-08-06 00:00:00").IsTomorrow() // true
	carbon.Parse("2020-08-06").IsTomorrow()          // true

	// 是否大于
	carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-04 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Gt(carbon.Parse("2020-08-05 13:14:15"))           // false
	carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-04 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").Compare(">", carbon.Parse("2020-08-05 13:14:15")) // false

	// 是否小于
	carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-06 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Lt(carbon.Parse("2020-08-05 13:14:15"))           // false
	carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-06 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").Compare("<", carbon.Parse("2020-08-05 13:14:15")) // false

	// 是否等于
	carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Eq(carbon.Parse("2020-08-05 13:14:00"))           // false
	carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").Compare("=", carbon.Parse("2020-08-05 13:14:00")) // false

	// 是否不等于
	carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-06 13:14:15"))            // true
	carbon.Parse("2020-08-05 13:14:15").Ne(carbon.Parse("2020-08-05 13:14:15"))            // false
	carbon.Parse("2020-08-05 13:14:15").Compare("!=", carbon.Parse("2020-08-06 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").Compare("<>", carbon.Parse("2020-08-05 13:14:15")) // false

	// 是否大于等于
	carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-04 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Gte(carbon.Parse("2020-08-05 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-04 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").Compare(">=", carbon.Parse("2020-08-05 13:14:15")) // true

	// 是否小于等于
	carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-06 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Lte(carbon.Parse("2020-08-05 13:14:15"))           // true
	carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-06 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").Compare("<=", carbon.Parse("2020-08-05 13:14:15")) // true

	// 是否在两个时间之间(不包括这两个时间)
	carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // false
	carbon.Parse("2020-08-05 13:14:15").Between(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

	// 是否在两个时间之间(包括开始时间)
	carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").BetweenIncludedStart(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

	// 是否在两个时间之间(包括结束时间)
	carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true

	// 是否在两个时间之间(包括这两个时间)
	carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-05 13:14:15"), carbon.Parse("2020-08-06 13:14:15")) // true
	carbon.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(carbon.Parse("2020-08-04 13:14:15"), carbon.Parse("2020-08-05 13:14:15")) // true

	// 设置时区
	carbon.SetTimezone(carbon.PRC).Now().ToDateTimeString()                           // 2020-08-05 13:14:15
	carbon.SetTimezone(carbon.Tokyo).Now().ToDateTimeString()                         // 2020-08-05 14:14:15
	carbon.SetTimezone(carbon.Tokyo).Now().SetTimezone(carbon.PRC).ToDateTimeString() // 2020-08-05 12:14:15

	// 设置区域
	carbon.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()    // 1 month ago
	carbon.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

	// 设置年月日时分秒
	carbon.Parse("2020-01-01").SetDateTime(2019, 2, 2, 13, 14, 15).ToString()  // 2019-02-02 13:14:15 +0800 CST
	carbon.Parse("2020-01-01").SetDateTime(2019, 2, 31, 13, 14, 15).ToString() // 2019-03-03 13:14:15 +0800 CST
	// 设置年月日时分秒毫秒
	carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 2, 13, 14, 15, 999).ToString()  // 2019-02-02 13:14:15.999 +0800 CST
	carbon.Parse("2020-01-01").SetDateTimeMilli(2019, 2, 31, 13, 14, 15, 999).ToString() // 2019-03-03 13:14:15.999 +0800 CST
	// 设置年月日时分秒微秒
	carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 2, 13, 14, 15, 999999).ToString()  // 2019-02-02 13:14:15.999999 +0800 CST
	carbon.Parse("2020-01-01").SetDateTimeMicro(2019, 2, 31, 13, 14, 15, 999999).ToString() // 2019-03-03 13:14:15.999999 +0800 CST
	// 设置年月日时分秒纳秒
	carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 2, 13, 14, 15, 999999999).ToString()  // 2019-02-02 13:14:15.999999999 +0800 CST
	carbon.Parse("2020-01-01").SetDateTimeNano(2019, 2, 31, 13, 14, 15, 999999999).ToString() // 2019-03-03 13:14:15.999999999 +0800 CST

	// 设置年月日
	carbon.Parse("2020-01-01").SetDate(2019, 2, 2).ToString()  // 2019-02-02 00:00:00 +0800 CST
	carbon.Parse("2020-01-01").SetDate(2019, 2, 31).ToString() // 2019-03-03 00:00:00 +0800 CST
	// 设置年月日毫秒
	carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 2, 999).ToString()  // 2019-02-02 00:00:00.999 +0800 CST
	carbon.Parse("2020-01-01").SetDateMilli(2019, 2, 31, 999).ToString() // 2019-03-03 00:00:00.999 +0800 CST
	// 设置年月日微秒
	carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 2, 999999).ToString()  // 2019-02-02 00:00:00.999999 +0800 CST
	carbon.Parse("2020-01-01").SetDateMicro(2019, 2, 31, 999999).ToString() // 2019-03-03 00:00:00.999999 +0800 CST
	// 设置年月日纳秒
	carbon.Parse("2020-01-01").SetDateNano(2019, 2, 2, 999999999).ToString()  // 2019-02-02 00:00:00.999999999 +0800 CST
	carbon.Parse("2020-01-01").SetDateNano(2019, 2, 31, 999999999).ToString() // 2019-03-03 00:00:00.999999999 +0800 CST

	// 设置时分秒
	carbon.Parse("2020-01-01").SetTime(13, 14, 15).ToString() // 2020-01-01 13:14:15 +0800 CST
	carbon.Parse("2020-01-01").SetTime(13, 14, 90).ToString() // 2020-01-01 13:15:30 +0800 CST
	// 设置时分秒毫秒
	carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 15, 999).ToString() // 2020-01-01 13:14:15.999 +0800 CST
	carbon.Parse("2020-01-01").SetTimeMilli(13, 14, 90, 999).ToString() // 2020-01-01 13:15:30.999 +0800 CST
	// 设置时分秒微秒
	carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 15, 999999).ToString() // 2020-01-01 13:14:15.999999 +0800 CST
	carbon.Parse("2020-01-01").SetTimeMicro(13, 14, 90, 999999).ToString() // 2020-01-01 13:15:30.999999 +0800 CST
	// 设置时分秒纳秒
	carbon.Parse("2020-01-01").SetTimeNano(13, 14, 15, 999999999).ToString() // 2020-01-01 13:14:15.999999999 +0800 CST
	carbon.Parse("2020-01-01").SetTimeNano(13, 14, 90, 999999999).ToString() // 2020-01-01 13:15:30.999999999 +0800 CST

	// 设置年份
	carbon.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
	// 设置年份(月份不溢出)
	carbon.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

	// 设置月份
	carbon.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
	// 设置月份(月份不溢出)
	carbon.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

	// 设置一周的开始日期
	carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
	carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0

	// 设置日期
	carbon.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
	carbon.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

	// 设置小时
	carbon.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
	carbon.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

	// 设置分钟
	carbon.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
	carbon.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

	// 设置秒
	carbon.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
	carbon.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

	// 设置毫秒
	carbon.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
	carbon.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

	// 设置微妙
	carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
	carbon.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

	// 设置纳秒
	carbon.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
	carbon.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999

	// 获取本年总天数
	carbon.Parse("2019-08-05 13:14:15").DaysInYear() // 365
	carbon.Parse("2020-08-05 13:14:15").DaysInYear() // 366
	// 获取本月总天数
	carbon.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
	carbon.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
	carbon.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

	// 获取本年第几天
	carbon.Parse("2020-08-05 13:14:15").DayOfYear() // 218
	// 获取本年第几周
	carbon.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
	carbon.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
	// 获取本月第几天
	carbon.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
	// 获取本月第几周
	carbon.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
	// 获取本周第几天
	carbon.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

	// 获取当前年月日时分秒
	carbon.Parse("2020-08-05 13:14:15").DateTime() // 2020,8,5,13,14,15
	// 获取当前年月日时分秒毫秒
	carbon.Parse("2020-08-05 13:14:15").DateTimeMilli() // 2020,8,5,13,14,15,999
	// 获取当前年月日时分秒微秒
	carbon.Parse("2020-08-05 13:14:15").DateTimeMicro() // 2020,8,5,13,14,15,999999
	// 获取当前年月日时分秒纳秒
	carbon.Parse("2020-08-05 13:14:15").DateTimeNano() // 2020,8,5,13,14,15,999999999

	// 获取当前年月日
	carbon.Parse("2020-08-05 13:14:15.999999999").Date() // 2020,8,5
	// 获取当前年月日毫秒
	carbon.Parse("2020-08-05 13:14:15.999999999").DateMilli() // 2020,8,5,999
	// 获取当前年月日微秒
	carbon.Parse("2020-08-05 13:14:15.999999999").DateMicro() // 2020,8,5,999999
	// 获取当前年月日纳秒
	carbon.Parse("2020-08-05 13:14:15.999999999").DateNano() // 2020,8,5,999999999

	// 获取当前时分秒
	carbon.Parse("2020-08-05 13:14:15.999999999").Time() // 13,14,15
	// 获取当前时分秒毫秒
	carbon.Parse("2020-08-05 13:14:15.999999999").TimeMilli() // 13,14,15,999
	// 获取当前时分秒微秒
	carbon.Parse("2020-08-05 13:14:15.999999999").TimeMicro() // 13,14,15,999999
	// 获取当前时分秒纳秒
	carbon.Parse("2020-08-05 13:14:15.999999999").TimeNano() // 13,14,15,999999999

	// 获取当前世纪
	carbon.Parse("2020-08-05 13:14:15").Century() // 21
	// 获取当前年代
	carbon.Parse("2019-08-05 13:14:15").Decade() // 10
	carbon.Parse("2021-08-05 13:14:15").Decade() // 20
	// 获取当前年份
	carbon.Parse("2020-08-05 13:14:15").Year() // 2020
	// 获取当前季度
	carbon.Parse("2020-08-05 13:14:15").Quarter() // 3
	// 获取当前月份
	carbon.Parse("2020-08-05 13:14:15").Month() // 8
	// 获取当前周(从0开始)
	carbon.Parse("2020-08-02 13:14:15").Week()                       // 0
	carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Sunday).Week() // 0
	carbon.Parse("2020-08-02").SetWeekStartsAt(carbon.Monday).Week() // 6
	// 获取当前天数
	carbon.Parse("2020-08-05 13:14:15").Day() // 5
	// 获取当前小时
	carbon.Parse("2020-08-05 13:14:15").Hour() // 13
	// 获取当前分钟
	carbon.Parse("2020-08-05 13:14:15").Minute() // 14
	// 获取当前秒钟
	carbon.Parse("2020-08-05 13:14:15").Second() // 15
	// 获取当前毫秒
	carbon.Parse("2020-08-05 13:14:15.999").Millisecond() // 999
	// 获取当前微秒
	carbon.Parse("2020-08-05 13:14:15.999").Microsecond() // 999000
	// 获取当前纳秒
	carbon.Parse("2020-08-05 13:14:15.999").Nanosecond() // 999000000

	// 获取秒级时间戳
	carbon.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
	// 获取毫秒级时间戳
	carbon.Parse("2020-08-05 13:14:15").TimestampMilli() // 1596604455000
	// 获取微秒级时间戳
	carbon.Parse("2020-08-05 13:14:15").TimestampMicro() // 1596604455000000
	// 获取纳秒级时间戳
	carbon.Parse("2020-08-05 13:14:15").TimestampNano() // 1596604455000000000

	// 获取时区
	carbon.SetTimezone(carbon.PRC).Timezone()   // CST
	carbon.SetTimezone(carbon.Tokyo).Timezone() // JST

	// 获取位置
	carbon.SetTimezone(carbon.PRC).Location()   // PRC
	carbon.SetTimezone(carbon.Tokyo).Location() // Asia/Tokyo

	// 获取距离UTC时区的偏移量，单位秒
	carbon.SetTimezone(carbon.PRC).Offset()   // 28800
	carbon.SetTimezone(carbon.Tokyo).Offset() // 32400

	// 获取当前区域
	carbon.Now().Locale()                    // en
	carbon.Now().SetLocale("zh-CN").Locale() // zh-CN

	// 获取当前星座
	carbon.Now().Constellation()                    // Leo
	carbon.Now().SetLocale("en").Constellation()    // Leo
	carbon.Now().SetLocale("zh-CN").Constellation() // 狮子座

	// 获取当前季节
	carbon.Now().Season()                    // Summer
	carbon.Now().SetLocale("en").Season()    // Summer
	carbon.Now().SetLocale("zh-CN").Season() // 夏季

	// 获取年龄
	carbon.Parse("2002-01-01 13:14:15").Age() // 17
	carbon.Parse("2002-12-31 13:14:15").Age() // 18

	// 输出日期时间字符串
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeString() // 2020-08-05 13:14:15
	// 输出日期时间字符串，包含毫秒
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMilliString() // 2020-08-05 13:14:15.999
	// 输出日期时间字符串，包含微秒
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeMicroString() // 2020-08-05 13:14:15.999999
	// 输出日期时间字符串，包含纳秒
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToDateTimeNanoString() // 2020-08-05 13:14:15.999999999

	// 输出简写日期时间字符串
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeString() // 20200805131415
	// 输出简写日期时间字符串，包含毫秒
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMilliString() // 20200805131415.999
	// 输出简写日期时间字符串，包含微秒
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeMicroString() // 20200805131415.999999
	// 输出简写日期时间字符串，包含纳秒
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToShortDateTimeNanoString() // 20200805131415.999999999

	// 输出日期字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToDateString() // 2020-08-05
	// 输出日期字符串，包含毫秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMilliString() // 2020-08-05.999
	// 输出日期字符串，包含微秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToDateMicroString() // 2020-08-05.999999
	// 输出日期字符串，包含纳秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToDateNanoString() // 2020-08-05.999999999

	// 输出简写日期字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateString() // 20200805
	// 输出简写日期字符串，包含毫秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMilliString() // 20200805.999
	// 输出简写日期字符串，包含微秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateMicroString() // 20200805.999999
	// 输出简写日期字符串，包含纳秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortDateNanoString() // 20200805.999999999

	// 输出时间字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
	// 输出时间字符串，包含毫秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMilliString() // 13:14:15.999
	// 输出时间字符串，包含微秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeMicroString() // 13:14:15.999999
	// 输出时间字符串，包含纳秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeNanoString() // 13:14:15.999999999

	// 输出简写时间字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeString() // 131415
	// 输出简写时间字符串，包含毫秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMilliString() // 131415.999
	// 输出简写时间字符串，包含微秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeMicroString() // 131415.999999
	// 输出简写时间字符串，包含纳秒
	carbon.Parse("2020-08-05 13:14:15.999999999").ToShortTimeNanoString() // 131415.999999999

	// 输出 ANSIC 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToANSICString() // Wed Aug  5 13:14:15 2020
	// 输出 Atom 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
	// 输出 UnixDate 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
	// 输出 RubyDate 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
	// 输出 Kitchen 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
	// 输出 Cookie 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
	// 输出 DayDateTime 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
	// 输出 RSS 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
	// 输出 W3C 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00

	// 输出 ISO8601 格式字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601String() // 2020-08-05T13:14:15+08:00
	// 输出 ISO8601Milli 格式字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MilliString() // 2020-08-05T13:14:15.999+08:00
	// 输出 ISO8601Micro 格式字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601MicroString() // 2020-08-05T13:14:15.999999+08:00
	// 输出 ISO8601Nano 格式字符串
	carbon.Parse("2020-08-05 13:14:15.999999999").ToIso8601NanoString() // 2020-08-05T13:14:15.999999999+08:00

	// 输出 RFC822 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
	// 输出 RFC822Z 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
	// 输出 RFC850 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
	// 输出 RFC1036 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
	// 输出 RFC1123 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
	// 输出 RFC1123Z 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
	// 输出 RFC2822 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
	// 输出 RFC7231 格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT

	// 输出 RFC3339 格式字符串
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339String() // 2020-08-05T13:14:15+08:00
	// 输出 RFC3339Milli 格式字符串
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MilliString() // 2020-08-05T13:14:15.999+08:00
	// 输出 RFC3339Micro 格式字符串
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339MicroString() // 2020-08-05T13:14:15.999999+08:00
	// 输出 RFC3339Nano 格式字符串
	carbon.Parse("2020-08-05T13:14:15.999999999+08:00").ToRfc3339NanoString() // 2020-08-05T13:14:15.999999999+08:00

	// 输出日期时间字符串
	_ = fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15

	// 输出"2006-01-02 15:04:05.999999999 -0700 MST"格式字符串
	carbon.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15.999999 +0800 CST

	// 输出指定布局的字符串,Layout()是ToLayoutString()的简写
	carbon.Parse("2020-08-05 13:14:15").Layout(carbon.ISO8601Layout)        // 2020-08-05T13:14:15+08:00
	carbon.Parse("2020-08-05 13:14:15").Layout("20060102150405")            // 20200805131415
	carbon.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒")     // 2020年08月05日 13时14分15秒
	carbon.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15

	// 输出指定格式的字符串,Format()是ToFormatString()的简写(如果使用的字母与格式化字符冲突时，请使用\符号转义该字符)
	carbon.Parse("2020-08-05 13:14:15").Format("YmdHis")                    // 20200805131415
	carbon.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒")             // 2020年08月05日 13时14分15秒
	carbon.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A")   // Wednesday 5th of August 2020 01:14:15 PM
	carbon.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15

	// 获取星座
	carbon.Parse("2020-08-05 13:14:15").Constellation() // Leo

	// 是否是白羊座
	carbon.Parse("2020-08-05 13:14:15").IsAries() // false
	// 是否是金牛座
	carbon.Parse("2020-08-05 13:14:15").IsTaurus() // false
	// 是否是双子座
	carbon.Parse("2020-08-05 13:14:15").IsGemini() // false
	// 是否是巨蟹座
	carbon.Parse("2020-08-05 13:14:15").IsCancer() // false
	// 是否是狮子座
	carbon.Parse("2020-08-05 13:14:15").IsLeo() // true
	// 是否是处女座
	carbon.Parse("2020-08-05 13:14:15").IsVirgo() // false
	// 是否是天秤座
	carbon.Parse("2020-08-05 13:14:15").IsLibra() // false
	// 是否是天蝎座
	carbon.Parse("2020-08-05 13:14:15").IsScorpio() // false
	// 是否是射手座
	carbon.Parse("2020-08-05 13:14:15").IsSagittarius() // false
	// 是否是摩羯座
	carbon.Parse("2020-08-05 13:14:15").IsCapricorn() // false
	// 是否是水瓶座
	carbon.Parse("2020-08-05 13:14:15").IsAquarius() // false
	// 是否是双鱼座
	carbon.Parse("2020-08-05 13:14:15").IsPisces() // false

	// 获取季节
	carbon.Parse("2020-08-05 13:14:15").Season() // Summer

	// 本季节开始时间
	carbon.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
	// 本季节结束时间
	carbon.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

	// 是否是春季
	carbon.Parse("2020-08-05 13:14:15").IsSpring() // false
	// 是否是夏季
	carbon.Parse("2020-08-05 13:14:15").IsSummer() // true
	// 是否是秋季
	carbon.Parse("2020-08-05 13:14:15").IsAutumn() // false
	// 是否是冬季
	carbon.Parse("2020-08-05 13:14:15").IsWinter() // false

	// 获取农历生肖
	carbon.Parse("2020-08-05 13:14:15").Lunar().Animal() // 鼠

	// 获取农历节日
	carbon.Parse("2021-02-12 13:14:15").Lunar().Festival() // 春节

	// 获取农历年月日时分秒
	carbon.Parse("2020-08-05 13:14:15").Lunar().DateTime() // 2020, 6, 16, 13, 14, 15
	// 获取农历年月日
	carbon.Parse("2020-08-05 13:14:15").Lunar().Date() // 2020, 6, 16
	// 获取农历时分秒
	carbon.Parse("2020-08-05 13:14:15").Lunar().Time() // 13, 14, 15

	// 获取农历年年份
	carbon.Parse("2020-08-05 13:14:15").Lunar().Year() // 2020
	// 获取农历月月份
	carbon.Parse("2020-08-05 13:14:15").Lunar().Month() // 6
	// 获取农历闰月月份
	carbon.Parse("2020-08-05 13:14:15").Lunar().LeapMonth() // 4
	// 获取农历日日期
	carbon.Parse("2020-08-05 13:14:15").Lunar().Day() // 16
	// 获取农历 YYYY-MM-DD HH::ii::ss 格式字符串
	_ = fmt.Sprintf("%s", carbon.Parse("2020-08-05 13:14:15").Lunar()) // 2020-06-16 13:14:15

	// 获取农历年字符串
	carbon.Parse("2020-08-05 13:14:15").Lunar().ToYearString() // 二零二零
	// 获取农历月字符串
	carbon.Parse("2020-08-05 13:14:15").Lunar().ToMonthString() // 六月
	// 获取农历天字符串
	carbon.Parse("2020-08-05 13:14:15").Lunar().ToDayString() // 十六
	// 获取农历日期字符串
	carbon.Parse("2020-08-05 13:14:15").Lunar().ToDateString() // 二零二零年六月十六

	// 是否是农历闰年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapYear() // true
	// 是否是农历闰月
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsLeapMonth() // false

	// 是否是鼠年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsRatYear() // true
	// 是否是牛年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsOxYear() // false
	// 是否是虎年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsTigerYear() // false
	// 是否是兔年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsRabbitYear() // false
	// 是否是龙年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsDragonYear() // false
	// 是否是蛇年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsSnakeYear() // false
	// 是否是马年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsHorseYear() // false
	// 是否是羊年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsGoatYear() // false
	// 是否是猴年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsMonkeyYear() // false
	// 是否是鸡年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsRoosterYear() // false
	// 是否是狗年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsDogYear() // false
	// 是否是猪年
	carbon.Parse("2020-08-05 13:14:15").Lunar().IsPigYear() // false

	// 获取农历时辰
	carbon.Parse("2020-02-05 21:00:00").Lunar().DoubleHour() // 亥时

	// 是否是子时
	carbon.Parse("2020-03-21 00:00:00").Lunar().IsFirstDoubleHour() // true
	// 是否是丑时
	carbon.Parse("2020-03-21 01:00:00").Lunar().IsSecondDoubleHour() // true
	// 是否是寅时
	carbon.Parse("2020-03-21 03:00:00").Lunar().IsThirdDoubleHour() // true
	// 是否是卯时
	carbon.Parse("2020-03-21 05:00:00").Lunar().IsFourthDoubleHour() // true
	// 是否是辰时
	carbon.Parse("2020-03-21 07:00:00").Lunar().IsFifthDoubleHour() // true
	// 是否是巳时
	carbon.Parse("2020-03-21 09:00:00").Lunar().IsSixthDoubleHour() // true
	// 是否是午时
	carbon.Parse("2020-03-21 11:00:00").Lunar().IsSeventhDoubleHour() // true
	// 是否是未时
	carbon.Parse("2020-03-21 13:00:00").Lunar().IsEighthDoubleHour() // true
	// 是否是申时
	carbon.Parse("2020-03-21 15:00:00").Lunar().IsNinthDoubleHour() // true
	// 是否是酉时
	carbon.Parse("2020-03-21 17:00:00").Lunar().IsTenthDoubleHour() // true
	// 是否是戌时
	carbon.Parse("2020-03-21 19:00:00").Lunar().IsEleventhDoubleHour() // true
	// 是否是亥时
	carbon.Parse("2020-03-21 21:00:00").Lunar().IsTwelfthDoubleHour() // true

	lang := carbon.NewLanguage()
	lang.SetLocale("zh-CN")

	c := carbon.SetLanguage(lang)
	if c.Error != nil {
		// 错误处理
		log.Fatal(c)
	}

	c.Now().AddHours(1).DiffForHumans()      // 1 小时后
	c.Now().AddHours(1).ToMonthString()      // 八月
	c.Now().AddHours(1).ToShortMonthString() // 8月
	c.Now().AddHours(1).ToWeekString()       // 星期二
	c.Now().AddHours(1).ToShortWeekString()  // 周二
	c.Now().AddHours(1).Constellation()      // 狮子座
	c.Now().AddHours(1).Season()             // 夏季
}
