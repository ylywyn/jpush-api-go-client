package jpushclient

import (
	"bytes"
	"encoding/json"
)

type AdminRequest struct {
	AppName        string `json:"app_name,string"`
	AndroidPackage string `json:"android_package,string"`
	GroupName      string `json:"group_name,string"`
}

func (c *Client) CreateApp(req *AdminRequest) (map[string]interface{}, error) {
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
