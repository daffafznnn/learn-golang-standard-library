package main

import(
	"fmt"
	"time"
)

func main() {
		var now time.Time = time.Now()
		fmt.Println(now.Local())

		var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
		fmt.Println(utc)
		fmt.Println(utc.Local())

		formatter := "2006-01-02 15:04:05"

		value := "2024-04-04 04:04:04"
		valueTime, err := time.Parse(formatter, value)
		if err != nil {
			fmt.Println("Error:",err.Error())
		} else {
				fmt.Println(valueTime)
		}
		
}
