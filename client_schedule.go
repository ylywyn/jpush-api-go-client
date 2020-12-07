package jpushclient

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func (c *Client) ScheduleCreateTask(req *Schedule) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/schedules"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) ScheduleGetList(page int) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/schedules"
	if page > 0 {
		link += "?page=" + strconv.Itoa(page)
	}
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) ScheduleView(id string) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/schedules/" + id
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
func (c *Client) ScheduleViewMsgId(id string) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/schedules/" + id + "/msg_ids"
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
func (c *Client) ScheduleUpdate(id string, req *Schedule) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/schedules/" + id
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("PUT", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) ScheduleDelete(id string) ([]byte, error) {
	link := c.pushUrl + "/v3/schedules/" + id
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
