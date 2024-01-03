package jpushclient

func (c *Client) QueryEffectFiles() (map[string]any, error) {
	link := c.pushUrl + "/v3/files"
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) QueryFile(fileId string) (map[string]any, error) {
	link := c.pushUrl + "/v3/files/" + fileId
	resp, err := c.request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c *Client) DeleteFile(fileId string) ([]byte, error) {
	link := c.pushUrl + "/v3/files/" + fileId
	resp, err := c.request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
