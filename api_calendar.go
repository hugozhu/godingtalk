package godingtalk

import "time"

type Event struct {
	OAPIResponse
	Location    string
	Summary     string
	Description string
}

type EventList struct {
	OAPIResponse
	HasMore   bool
	EventList []Event
}

type CalendarTime struct {
	TimeZone string `json:"time_zone"`
	Date     string `json:"date"`
}

type CalendarRequest struct {
	TimeMax CalendarTime `json:"time_max"`
	TimeMin CalendarTime `json:"time_min"`
	StaffId string       `json:"staff_id"`
}

func (c *DingTalkClient) ListEvents(staffid string, from time.Time, to time.Time) (events EventList, err error) {
	location := time.Now().Location().String()
	timeMin := CalendarTime{
		TimeZone: location,
		Date:     from.Format("2006-01-02"),
	}
	timeMax := CalendarTime{
		TimeZone: location,
		Date:     to.Format("2006-01-02"),
	}

	data := map[string]CalendarRequest{
		"open_calendar_list_request": CalendarRequest{
			TimeMax: timeMax,
			TimeMin: timeMin,
			StaffId: staffid,
		},
	}
	err = c.httpRPC("topapi/calendar/list", nil, data, &events)
	return events, err
}
