package jpushclient

func (c *Client) FileGetList() (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/files"
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) FileView(fileId string) (map[string]interface{}, error) {
	link := c.pushUrl + "/v3/files/" + fileId
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) FileDeleteTag(fileId string) ([]byte, error) {
	link := c.pushUrl + "/v3/files/" + fileId
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
