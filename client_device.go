package jpushclient

import (
	"bytes"
	"encoding/json"
	"strings"
)

func (c *Client) DeviceView(registrationId string) (map[string]interface{}, error) {
	link := c.deviceUrl + "/v3/devices/" + registrationId
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) DeviceRequest(registrationId string, req *DeviceSettingRequest) ([]byte, error) {
	link := c.deviceUrl + "/v3/devices/" + registrationId
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

func (c *Client) DeviceEmptyTagsRequest(registrationId string, req *DeviceSettingEmptyTagsRequest) ([]byte, error) {
	link := c.deviceUrl + "/v3/devices/" + registrationId
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

//vip
func (c *Client) DeviceGetUserStatus(req []string) (map[string]interface{}, error) {
	link := c.deviceUrl + "/v3/devices/status/"
	params := make(map[string]interface{})
	params["registration_ids"] = req
	buf, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
func (c *Client) DeviceGetWithAlias(alias string, platforms []string) (map[string]interface{}, error) {
	link := c.deviceUrl + "/v3/aliases/" + alias
	if len(platforms) > 0 {
		link += "?platform=" + strings.Join(platforms, ",")
	}
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
func (c *Client) DeviceRemoveBindAlias(alias string, req *DeviceSettingRequestAlias) ([]byte, error) {
	link := c.deviceUrl + "/v3/aliases/" + alias
	params := make(map[string]interface{})
	params["registration_ids"] = req
	buf, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
func (c *Client) DeviceDeleteAlias(alias string) ([]byte, error) {
	link := c.deviceUrl + "/v3/aliases/" + alias
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

func (c *Client) DeviceGetTags() (map[string]interface{}, error) {
	link := c.deviceUrl + "/v3/tags/"
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) DeviceCheckDeviceWithTag(tag, registrationId string) (map[string]interface{}, error) {
	link := c.deviceUrl + "/v3/tags/" + tag + "/registration_ids/" + registrationId
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) DeviceBindTags(tag string, req *DeviceBindTagsRequest) ([]byte, error) {
	link := c.deviceUrl + "/v3/tags/" + tag
	params := make(map[string]interface{})
	params["registration_ids"] = req
	buf, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := c.request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

func (c *Client) DeviceDeleteTag(tag string, platforms []string) ([]byte, error) {
	link := c.deviceUrl + "/v3/tags/" + tag
	if len(platforms) > 0 {
		link += "?platform=" + strings.Join(platforms, ",")
	}
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
