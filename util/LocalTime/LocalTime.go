package LocalTime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// LocalTime 别名
type LocalTime time.Time

// 重写MarshalJSON方法来实现时间格式化
// 2006-01-02 15:04:05
func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// 字符串转时间对象
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}

// 写入数据库时会调用该方法将自定义时间类型转换并写入数据库
// 将LocalTime转为time.Time类型
func (t LocalTime) Value() (driver.Value, error) {
	// LocalTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

// 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
