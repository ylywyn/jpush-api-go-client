package jpushclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

func (c *Client) ReportReceived(msgIds []string) ([]interface{}, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("msgIds不能为空")
	}
	link := c.reportUrl + "/v3/received?msg_ids=" + strings.Join(msgIds, ",")
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}

//vip
func (c *Client) ReportReceivedDetail(msgIds []string) ([]interface{}, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("msgIds不能为空")
	}
	link := c.reportUrl + "/v3/received/detail?msg_ids=" + strings.Join(msgIds, ",")
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}

func (c *Client) ReportStatusMessage(req *ReportStatusRequest) (map[string]interface{}, error) {
	link := c.reportUrl + "/v3/status/message"
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

//vip
func (c *Client) ReportMessages(msgIds []string) ([]interface{}, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("msgIds不能为空")
	}
	link := c.reportUrl + "/v3/messages?msg_ids=" + strings.Join(msgIds, ",")
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}

//vip
func (c *Client) ReportMessagesDetail(msgIds []string) ([]interface{}, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("msgIds不能为空")
	}
	link := c.reportUrl + "/v3/messages/detail?msg_ids=" + strings.Join(msgIds, ",")
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}

func (c *Client) ReportGetUsers(timeUnit, start, duration string) ([]interface{}, error) {
	link := c.reportUrl + "/v3/users?time_unit=" + timeUnit + "&start" + start + "&duration" + duration
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}
func (c *Client) ReportGetGroupMessages(msgIds []string) ([]interface{}, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("msgIds不能为空")
	}
	link := c.reportUrl + "/v3/group/messages/detail?group_msgids=" + strings.Join(msgIds, ",")
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}
func (c *Client) ReportGetGroup(timeUnit, start, duration string) ([]interface{}, error) {
	link := c.reportUrl + "/v3/group/users/time_unit=" + timeUnit + "&start" + start + "&duration" + duration
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}
