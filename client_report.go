package jpushclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
)

func (c *Client) GetReceived(msgIds []string) ([]interface{}, error) {
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

/*
   送达统计详情（新）
   https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/#_7
*/
func (c *Client) GetReceivedDetail(msgIds []string) ([]interface{}, error) {
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

func (c *Client) GetMessageStatus(req *ReportStatusRequest) (map[string]interface{}, error) {
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

/*
	消息统计（VIP 专属接口，旧）
	https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/#vip
*/
func (c *Client) GetMessages(msgIds []string) ([]interface{}, error) {
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

/*
   消息统计详情（VIP 专属接口，新）
   https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/#vip_1
*/
func (c *Client) GetMessagesDetail(msgIds []string) ([]interface{}, error) {
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

func (c *Client) GetUsers(timeUnit, start, duration string) ([]interface{}, error) {
	link := c.reportUrl + "/v3/users?time_unit=" + timeUnit + "&start" + start + "&duration" + duration
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}

/*
	分组统计-消息统计（VIP 专属接口）
	https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/#-vip
*/
func (c *Client) GetGroupMessages(msgIds []string) ([]interface{}, error) {
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

/*
	分组统计-用户统计（VIP 专属接口）
	https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report/#-vip_1
*/
func (c *Client) GetGroupUsers(timeUnit, start, duration string) ([]interface{}, error) {
	link := c.reportUrl + "/v3/group/users/time_unit=" + timeUnit + "&start" + start + "&duration" + duration
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Array()
}
