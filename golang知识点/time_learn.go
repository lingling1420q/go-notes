//time库的用法
package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	currentTime := time.Now()
	fmt.Println(currentTime)
	// 年月日的获取,返回的是时间
	Y := currentTime.Year()
	fmt.Println(Y)
	fmt.Printf("%T", Y)

	// 通过Date来获取当前时间
	currentTimeDate := time.Date(2019, 5, 7, 18, 59, 59, 0, time.Local)
	fmt.Println(currentTimeDate)

	// 当前时间戳
	//单位s,打印结果:1491888244
	timeUnix := time.Now().Unix()
	//单位纳秒,打印结果：1491888244752784461
	timeUnixNano := time.Now().UnixNano()
	fmt.Println(timeUnix, timeUnixNano)

	//当前时间的字符串，2006-01-02 15:04:05据说是golang的诞生时间，固定写法
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)

	// 时间戳转时间字符串 (int64 —>  string)
	timeUnix = time.Now().Unix()
	timeToStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
	fmt.Println(timeToStr)

	// 时间字符串转时间(string => Time)
	formatTimeStr := "2017-04-11 13:33:37"
	formatTime, err := time.Parse("2006-01-02 15:04:05", formatTimeStr)
	if err == nil {
		fmt.Println(formatTime)
	}
	// 时间字符串转时间戳
	fmt.Println(formatTime.Unix())
}
