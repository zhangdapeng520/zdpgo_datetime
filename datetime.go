package zdpgo_datetime

import "github.com/zhangdapeng520/zdpgo_datetime/carbon"

/*
@Time : 2022/5/19 22:56
@Author : 张大鹏
@File : datetime
@Software: Goland2021.3.1
@Description: datetime 日期时间核心方法相关
*/

type DateTime struct {
	Config *Config // 配置

	TimeNow       carbon.Carbon // 今天此刻
	TimeYesterday carbon.Carbon // 昨天此刻
	TimeTomorrow  carbon.Carbon // 明天此刻

	Time carbon.Carbon // 时间对象
}

func New() *DateTime {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *DateTime {
	d := &DateTime{}
	d.Config = config

	// 初始化时间
	d.TimeNow = carbon.Now(Shanghai)
	d.TimeYesterday = carbon.Yesterday(Shanghai)
	d.TimeTomorrow = carbon.Tomorrow(Shanghai)

	return d
}
