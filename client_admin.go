package jpushclient

import (
	"bytes"
	"encoding/json"
)

func (c *Client) CreateApp(req *AdminRequest) (map[string]any, error) {
	link := c.adminUrl + "/v1/app"
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

func (c *Client) DeleteApp(appKey string) ([]byte, error) {
	link := c.pushUrl + "/v1/app/" + appKey + "/delete"
	resp, err := c.request("POST", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
