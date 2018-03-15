package godingtalk

import (
	"testing"
	"time"
)

func TestListAttendanceRecord(t *testing.T) {
	dataFrom, _ := time.Parse("2006-01-02", "2018-03-06")
	dataTo, _ := time.Parse("2006-01-02", "2018-03-10")
	records, err := c.ListAttendanceRecord([]string{"085354234826136236"}, dataFrom, dataTo)
	if err != nil {
		t.Error(err)
	} else if len(records) > 0 {
		t.Logf("%+v\n", records)
	}
}

func TestListAttendanceResult(t *testing.T) {
	dataFrom, _ := time.Parse("2006-01-02", "2018-03-06")
	dataTo, _ := time.Parse("2006-01-02", "2018-03-10")
	resp, err := c.ListAttendanceResult([]string{"085354234826136236"}, dataFrom, dataTo, 0, 2)
	if err != nil {
		t.Error(err)
	} else if len(resp.Records) > 0 {
		t.Logf("%+v\n", resp.Records[0])
	}
}
