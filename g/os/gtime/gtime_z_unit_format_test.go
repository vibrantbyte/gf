package gtime_test

import (
	"github.com/gogf/gf/g/os/gtime"
	"github.com/gogf/gf/g/test/gtest"
	"testing"
)

func Test_Format(t *testing.T) {
	gtest.Case(t, func() {
		timeTemp, err := gtime.StrToTime("2006-01-11 15:04:05", "Y-m-d H:i:s")
		timeTemp.ToZone("Asia/Shanghai")
		if err != nil {
			t.Error("test fail")
		}
		gtest.Assert(timeTemp.Format("\\T\\i\\m\\e中文Y-m-j G:i:s.u\\"), "Time中文2006-01-11 15:04:05.000")

		gtest.Assert(timeTemp.Format("d D j l"), "11 Wed 11 Wednesday")

		gtest.Assert(timeTemp.Format("F m M n"), "January 01 Jan 1")

		gtest.Assert(timeTemp.Format("Y y"), "2006 06")

		gtest.Assert(timeTemp.Format("a A g G h H i s u .u"), "pm PM 3 15 03 15 04 05 000 .000")

		gtest.Assert(timeTemp.Format("O P T"), "+0800 +08:00 CST")

		gtest.Assert(timeTemp.Format("r"), "Wed, 11 Jan 06 15:04 CST")

		gtest.Assert(timeTemp.Format("c"), "2006-01-11T15:04:05+08:00")

		//补零
		timeTemp1, err := gtime.StrToTime("2006-01-02 03:04:05", "Y-m-d H:i:s")
		if err != nil {
			t.Error("test fail")
		}
		gtest.Assert(timeTemp1.Format("Y-m-d h:i:s"), "2006-01-02 03:04:05")
		//不补零
		timeTemp2, err := gtime.StrToTime("2006-01-02 03:04:05", "Y-m-d H:i:s")
		if err != nil {
			t.Error("test fail")
		}
		gtest.Assert(timeTemp2.Format("Y-n-j G:i:s"), "2006-1-2 3:04:05")

		// 测试数字型的星期
		times := []map[string]string{
			{"k": "2019-04-22", "f": "w", "r": "1"},
			{"k": "2019-04-27", "f": "w", "r": "6"},
			{"k": "2019-03-10", "f": "w", "r": "0"},
			{"k": "2019-03-10", "f": "Y-m-d 星期:w", "r": "2019-03-10 星期:0"},
		}

		for _, v := range times {
			t1, err1 := gtime.StrToTime(v["k"], "Y-m-d")
			gtest.Assert(err1, nil)
			gtest.Assert(t1.Format(v["f"]), v["r"])
		}

	})
}

func Test_Layout(t *testing.T) {
	gtest.Case(t, func() {
		timeTemp := gtime.Now()
		gtest.Assert(timeTemp.Layout("2006-01-02 15:04:05"), timeTemp.Time.Format("2006-01-02 15:04:05"))
	})
}
