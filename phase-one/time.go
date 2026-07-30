package main

import (
	"fmt"
	"time"
)

const (
	TIME_FMT = "2006-01-02 15:04:05.000"
	DATE_FMT = "20060102"
)

func main25() {
	fmt.Println("start")
	time.Sleep(3 * time.Second)
	fmt.Println("bye-bye")

	t0 := time.Now()
	fmt.Printf("时间戳 秒%d, 毫秒%d, 微秒%d, 纳秒%d\n", t0.Unix(), t0.UnixMilli(), t0.UnixMicro(), t0.UnixNano())
	time.Sleep(3 * time.Second)
	t1 := time.Now()
	diff1 := t1.Sub(t0)
	fmt.Printf("时间差 %f秒\n", diff1.Seconds())
	diff2 := time.Since(t0)
	fmt.Printf("时间差 %f秒\n", diff2.Seconds())
	fmt.Printf("t1>t0吗？ %t\n", t1.After(t0))
	fmt.Printf("t1>t0吗？ %t\n", t1.Before(t0))

	d := time.Duration(3 * time.Hour)
	t3 := t0.Add(d)
	fmt.Printf("%d-%d-%d %d:%d:%d.%d\n", t3.Year(), t3.Month(), t3.Day(), t3.Hour(), t3.Minute(), t3.Second(), t3.UnixMilli()%t3.Unix())

	fmt.Printf("week day %s %d\n", t3.Weekday().String(), t3.Weekday()) // 周日的Weekday()是0，周六的Weekday()是6
	fmt.Printf("day in year %d\n", t3.YearDay())                        // 属于一年当中的第几天
	fmt.Printf("t3= %s\n", t3.Format(TIME_FMT))
	fmt.Printf("t3= %s\n", t3.Format(DATE_FMT))

	ts := t3.Format(TIME_FMT)

	t4, err := time.Parse(TIME_FMT, ts) // 不建议使用Parse
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("t4 = %s\n", t4.Format(TIME_FMT))
	}

	loc, _ := time.LoadLocation("Asia/Shanghai")
	t5, err := time.ParseInLocation(TIME_FMT, ts, loc) // 生产环境一律使用ParseInLocation
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("t5 = %s\n", t5.Format(TIME_FMT))
	}
}
