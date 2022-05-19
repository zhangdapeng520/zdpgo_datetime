package zdpgo_datetime

import "github.com/zhangdapeng520/zdpgo_datetime/carbon"

/*
@Time : 2022/5/19 23:19
@Author : 张大鹏
@File : set
@Software: Goland2021.3.1
@Description: set设置相关
*/

// SetTimezone 设置时区
func (d *DateTime) SetTimezone(timezone string) {
	d.TimeNow = carbon.Now(timezone)
}
