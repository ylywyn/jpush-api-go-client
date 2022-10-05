package jpushclient

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func (c *Client) Push(push *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) GetCidPool(count int, cidType string) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/cid?"
	if count > 0 {
		link += "count=" + strconv.Itoa(count)
	}
	if cidType != "" {
		link += "type=" + cidType
	}
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) GroupPush(push *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/grouppush"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), true)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) Validate(req *PushRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/validate"
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

func (c *Client) BatchPushByRegId(push *PushSingleRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/batch/regid/single"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) BatchPushByAlias(push *PushSingleRequest) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/push/batch/alias/single"
	buf, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) DeletePush(msgId string) ([]byte, error) {
	link := c.pushUrl + "/v3/push/" + msgId
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
