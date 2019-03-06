package godingtalk

import "time"

type Event struct {
	OAPIResponse
	Id          int64
	UUID        string `json:"unique_id"`
	Location    string
	Summary     string
	Description string
	Start       struct {
		DateTime string `json:"date_time"`
	}
	End struct {
		DateTime string `json:"date_time"`
	}
}

type ListEventsResponse struct {
	OAPIResponse
	Result struct {
		OAPIResponse
		Data struct {
			Events []Event `json:"items"`
		} `json:"result"`
	} `json:"result"`
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

func (c *DingTalkClient) ListEvents(staffid string, from time.Time, to time.Time) (events []Event, err error) {
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
	var resp ListEventsResponse
	err = c.httpRPC("topapi/calendar/list", nil, data, &resp)
	events = resp.Result.Data.Events
	return events, err
}
