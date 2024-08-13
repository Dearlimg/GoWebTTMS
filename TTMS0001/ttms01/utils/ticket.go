package utils

import (
	"strconv"
	"strings"
)

func PackTicketData(nums []int) string {
	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
		if i != len(nums)-1 {
			res += "_"
		}
	}
	return res
}

func ParseTicketData(data string) string {
	str := strings.Replace(data, "_", " ", -1)
	return str
}

func SeatsToNumbers(seats string) string {
	var seatNumbers []string
	for i, seat := range seats {
		if seat == '1' {
			// 座位号从1开始，所以i+1
			seatNumbers = append(seatNumbers, strconv.Itoa(i+1))
		}
	}
	// 使用下划线连接所有座位号
	return strings.Join(seatNumbers, "_")
}
