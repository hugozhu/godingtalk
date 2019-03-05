package godingtalk

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

func (c *DingTalkClient) ListEvents() (events EventList, err error) {
	timeMin := CalendarTime{
		TimeZone: "Asia/Shanghai",
		Date:     "2019-02-03",
	}
	timeMax := CalendarTime{
		TimeZone: "Asia/Shanghai",
		Date:     "2019-03-04",
	}

	data := map[string]CalendarRequest{
		"open_calendar_list_request": CalendarRequest{
			TimeMax: timeMax,
			TimeMin: timeMin,
			StaffId: "0420506555",
		},
	}
	err = c.httpRPC("topapi/calendar/list", nil, data, &events)
	return events, err
}
