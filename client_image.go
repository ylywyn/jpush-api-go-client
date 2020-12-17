package jpushclient

import (
	"bytes"
	"encoding/json"
)

func (c *Client) AddImageUrl(req *ImageRequest) (map[string]interface{}, error) {
	link := c.imageUrl + "/byurls"
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

func (c *Client) UpdateImageUrl(mediaId string, req *ImageRequest) (map[string]interface{}, error) {
	link := c.imageUrl + "/byurls/" + mediaId
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
