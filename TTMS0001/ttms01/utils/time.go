package utils

import (
	"strconv"
	"time"
)

func isTimeOverdue(startTimeStr, endTimeStr, minutesStr string) (bool, error) {
	// 定义时间格式
	timeFormat := "2006-01-02-15:04"

	// 解析时间字符串
	startTime, err := time.Parse(timeFormat, startTimeStr)
	if err != nil {
		return false, err
	}
	endTime, err := time.Parse(timeFormat, endTimeStr)
	if err != nil {
		return false, err
	}

	// 转换分钟数
	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return false, err
	}

	// 计算经过的时间
	dueTime := startTime.Add(time.Duration(minutes) * time.Minute)

	// 判断是否超时
	return dueTime.After(endTime), nil
}

func IsWithinRange(time1Str, time2Str, minutesStr string) (bool, error) {
	// 时间格式
	timeFormat := "2006-01-02-15:04"

	// 解析时间字符串为 time.Time 对象
	time1, err := time.Parse(timeFormat, time1Str)
	if err != nil {
		return false, err
	}
	time2, err := time.Parse(timeFormat, time2Str)
	if err != nil {
		return false, err
	}

	// 解析分钟数
	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return false, err
	}

	// 计算时间范围
	timeRangeStart := time2.Add(-time.Duration(minutes) * time.Minute)
	timeRangeEnd := time2.Add(time.Duration(minutes) * time.Minute)

	// 判断 time1 是否在 time2 的范围内
	return time1.After(timeRangeStart) && time1.Before(timeRangeEnd), nil
}
