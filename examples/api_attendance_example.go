package main

import (
	"fmt"
	"github.com/ipandtcp/godingtalk"
	"os"
	"time"
)

func main() {

	c := godingtalk.NewDingTalkClient(os.Getenv("corpid"), os.Getenv("corpsecret"))
	err := c.RefreshAccessToken()
	if err != nil {
		panic(err)
	}

	dataFrom, _ := time.Parse("2006-01-02", "2018-03-06")
	dataTo, _ := time.Parse("2006-01-02", "2018-03-10")
	record, err := c.ListAttendanceRecord([]string{"085354234826136236"}, dataFrom, dataTo)
	if err != nil {
		panic(err)
	} else if len(record.Records) > 0 {
		fmt.Printf("%#v\n", record.Records[0])
	}

	result, err := c.ListAttendanceResult([]string{"085354234826136236"}, dataFrom, dataTo, 0, 2)
	if err != nil {
		panic(err)
	} else if len(result.Records) > 0 {
		fmt.Printf("%#v\n", result.Records[0])
	}
}
