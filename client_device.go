package jpushclient

import (
	"bytes"
	"encoding/json"
	"strings"
)

func (c *Client) GetDevices(registrationId string) (map[string]any, error) {
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

// GetDevicesStatus by vip
func (c *Client) GetDevicesStatus(req []string) (map[string]any, error) {
	link := c.deviceUrl + "/v3/devices/status/"
	params := make(map[string]any)
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

func (c *Client) GetAliasDevices(alias string, platforms []string) (map[string]any, error) {
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

func (c *Client) RemoveAlias(alias string, req *DeviceSettingRequestAlias) ([]byte, error) {
	link := c.deviceUrl + "/v3/aliases/" + alias
	params := make(map[string]any)
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

func (c *Client) DeleteAlias(alias string) ([]byte, error) {
	link := c.deviceUrl + "/v3/aliases/" + alias
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}

func (c *Client) GetTags() (map[string]any, error) {
	link := c.deviceUrl + "/v3/tags/"
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) IsDeviceInTag(tag, registrationId string) (map[string]any, error) {
	link := c.deviceUrl + "/v3/tags/" + tag + "/registration_ids/" + registrationId
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) UpdateTag(tag string, req *DeviceBindTagsRequest) ([]byte, error) {
	link := c.deviceUrl + "/v3/tags/" + tag
	params := make(map[string]any)
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

func (c *Client) DeleteTag(tag string, platforms []string) ([]byte, error) {
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
