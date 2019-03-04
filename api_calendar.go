package godingtalk

import "net/url"

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

func (c *DingTalkClient) ListEvents() (events EventList, err error) {
	params := url.Values{}
	err = c.httpRPC("topapi/calendar/list", params, nil, &events)
	return events, err
}
